package k8s

import (
	"context"
	"fmt"

	"github.com/ez-deploy/ezdeploy/models"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

var (
	ERRServiceNotFound = fmt.Errorf("service not found")
	ERRServiceConflict = fmt.Errorf("service already exists")
)

// ServiceManager impl service operations.
type ServiceManager struct {
	ClientSet *kubernetes.Clientset
}

// SetService create service if not exist, update if exist.
func (m *ServiceManager) SetService(
	ctx context.Context,
	namespace string,
	service *models.ServiceInfo,
	versionInfo *models.ServiceVersion,
) error {
	servicesClient := m.ClientSet.CoreV1().Services(namespace)
	_, err := servicesClient.Get(ctx, service.Name, metav1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return err
	}

	// create service if not exist.
	if errors.IsNotFound(err) {
		return m.createService(ctx, namespace, service, versionInfo)
	}

	return m.updateService(ctx, namespace, service, versionInfo)
}

func (m *ServiceManager) DeleteService(
	ctx context.Context,
	namespace string,
	service string,
) error {
	servicesClient := m.ClientSet.CoreV1().Services(namespace)
	return servicesClient.Delete(ctx, service, metav1.DeleteOptions{})
}

func (m *ServiceManager) createService(
	ctx context.Context,
	namespace string,
	service *models.ServiceInfo,
	versionInfo *models.ServiceVersion,
) error {
	// create service in k8s.
	servicesClient := m.ClientSet.CoreV1().Services(namespace)
	svc := buildServiceConfigFromServiceInfo(service, versionInfo)

	_, err := servicesClient.Create(ctx, svc, metav1.CreateOptions{})
	if errors.IsAlreadyExists(err) {
		return ERRServiceConflict
	}

	return err
}

func (m *ServiceManager) updateService(
	ctx context.Context,
	namespace string,
	service *models.ServiceInfo,
	versionInfo *models.ServiceVersion,
) error {
	// update service in k8s.
	servicesClient := m.ClientSet.CoreV1().Services(namespace)
	svc := buildServiceConfigFromServiceInfo(service, versionInfo)

	_, err := servicesClient.Update(ctx, svc, metav1.UpdateOptions{})
	if errors.IsNotFound(err) {
		return ERRNamespaceNotExist
	}

	return err
}

func buildServiceConfigFromServiceInfo(
	service *models.ServiceInfo,
	versionInfo *models.ServiceVersion) *apiv1.Service {

	var (
		name = service.Name

		// service type.
		exposeType     = apiv1.ServiceTypeClusterIP
		containerPort  = (int)(versionInfo.ContainerPort)
		inClusterPort  = (int32)(service.InClusterPort)
		outClusterPort = (int32)(service.NodePort)
	)

	if service.ExposeType == models.ServiceInfoExposeTypeNodeport {
		exposeType = apiv1.ServiceTypeNodePort
	}

	svc := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Labels: map[string]string{
				selectorName: name,
			},
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				selectorName: name,
			},
			Type: exposeType,
			Ports: []apiv1.ServicePort{
				{
					Name:       name,
					Protocol:   apiv1.ProtocolTCP,
					TargetPort: intstr.FromInt(containerPort),
					Port:       inClusterPort,
					NodePort:   outClusterPort,
				},
			},
		},
	}

	return svc
}
