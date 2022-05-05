package k8s

import (
	"context"
	"fmt"

	"github.com/ez-deploy/ezdeploy/models"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
)

const selectorName = "app-name"

var (
	ERRDeploymentConflict = fmt.Errorf("deployment already exists")
	ERRDeploymentNotFound = fmt.Errorf("deployment not found")
)

// DeploymentManager impl deployment operations.
type DeploymentManager struct {
	ClientSet *kubernetes.Clientset
}

// SetDeployment create deployment if not exist, update if exist.
func (m *DeploymentManager) SetDeployment(
	ctx context.Context,
	namespace string,
	service *models.ServiceInfo,
	versionInfo *models.ServiceVersion,
) error {
	deploymentsClient := m.ClientSet.AppsV1().Deployments(namespace)

	// check if deployment exist.
	_, err := deploymentsClient.Get(ctx, service.Name, metav1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	// create deployment if not exist.
	if errors.IsNotFound(err) {
		return m.createDeployment(ctx, namespace, service, versionInfo)
	}

	// update deployment if exist.
	return m.updateDeployment(ctx, namespace, service, versionInfo)
}

func (m *DeploymentManager) createDeployment(
	ctx context.Context,
	namespace string,
	service *models.ServiceInfo,
	versionInfo *models.ServiceVersion,
) error {

	deploymentsClient := m.ClientSet.AppsV1().Deployments(namespace)

	deploymentConfig := buildDeploymentConfigFromServiceInfo(service, versionInfo)

	_, err := deploymentsClient.Create(ctx, deploymentConfig, metav1.CreateOptions{})
	if errors.IsAlreadyExists(err) {
		return ERRDeploymentConflict
	}

	return err
}

func (m *DeploymentManager) updateDeployment(
	ctx context.Context,
	namespace string,
	service *models.ServiceInfo,
	versionInfo *models.ServiceVersion,
) error {
	// update service in k8s.
	deploymentsClient := m.ClientSet.AppsV1().Deployments(namespace)
	expectConfig := buildDeploymentConfigFromServiceInfo(service, versionInfo)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentsClient.Get(ctx, service.Name, metav1.GetOptions{})
		if getErr != nil {
			return fmt.Errorf("failed to get latest version of Service: %v", getErr)
		}

		result.Spec.Replicas = expectConfig.Spec.Replicas
		result.Spec.Template = expectConfig.Spec.Template

		_, updateErr := deploymentsClient.Update(ctx, result, metav1.UpdateOptions{})
		return updateErr
	})

	return retryErr
}

// DeleteDeployment delete deployment.
func (m *DeploymentManager) DeleteDeployment(ctx context.Context, namespace, name string) error {
	deploymentsClient := m.ClientSet.AppsV1().Deployments(namespace)

	deletePolicy := metav1.DeletePropagationForeground
	deleteOpt := metav1.DeleteOptions{PropagationPolicy: &deletePolicy}

	return deploymentsClient.Delete(ctx, name, deleteOpt)
}

func buildDeploymentConfigFromServiceInfo(
	service *models.ServiceInfo,
	versionInfo *models.ServiceVersion,
) *appsv1.Deployment {
	var (
		name    = service.Name
		replica = (int32)(service.Replica)

		imageURL   = versionInfo.Image
		exposePort = (int32)(versionInfo.ContainerPort)
	)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replica,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					selectorName: name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						selectorName: name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  name,
							Image: imageURL,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "tcp",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: exposePort,
								},
							},
						},
					},
				},
			},
		},
	}

	return deployment
}
