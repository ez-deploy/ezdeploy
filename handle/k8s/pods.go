package k8s

import (
	"context"
	"fmt"

	"github.com/ez-deploy/ezdeploy/models"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodsManager struct {
	ClientSet *kubernetes.Clientset
}

func (m *PodsManager) ListPods(ctx context.Context, namespace, service string) ([]*models.PodInfo, error) {
	podsClient := m.ClientSet.CoreV1().Pods(namespace)
	listOpt := v1.ListOptions{LabelSelector: fmt.Sprintf("%s=%s", selectorName, service)}
	pods, err := podsClient.List(ctx, listOpt)
	if err != nil {
		return nil, err
	}

	var podsInfo []*models.PodInfo
	for _, pod := range pods.Items {
		podsInfo = append(podsInfo, &models.PodInfo{
			Name:   pod.Name,
			IP:     pod.Status.PodIP,
			Status: string(pod.Status.Phase),
		})
	}

	return podsInfo, nil
}
