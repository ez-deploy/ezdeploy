package k8s

import "k8s.io/client-go/kubernetes"

// DeploymentManager impl deployment operations.
type DeploymentManager struct {
	ClientSet *kubernetes.Clientset
}
