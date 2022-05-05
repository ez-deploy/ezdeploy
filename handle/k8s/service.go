package k8s

import "k8s.io/client-go/kubernetes"

// ServiceManager impl service operations.
type ServiceManager struct {
	ClientSet *kubernetes.Clientset
}
