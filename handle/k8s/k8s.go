package k8s

import "k8s.io/client-go/kubernetes"

type K8SManager struct {
	*NamespaceManager
}

func New(clientSet *kubernetes.Clientset) *K8SManager {
	return &K8SManager{
		NamespaceManager: &NamespaceManager{ClientSet: clientSet},
	}
}
