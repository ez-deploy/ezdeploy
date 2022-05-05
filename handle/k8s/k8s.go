package k8s

import "k8s.io/client-go/kubernetes"

type K8SManager struct {
	*NamespaceManager
	*ServiceManager
	*DeploymentManager
	*PodsManager
}

func New(clientSet *kubernetes.Clientset) *K8SManager {
	return &K8SManager{
		NamespaceManager:  &NamespaceManager{ClientSet: clientSet},
		ServiceManager:    &ServiceManager{ClientSet: clientSet},
		DeploymentManager: &DeploymentManager{ClientSet: clientSet},
		PodsManager:       &PodsManager{ClientSet: clientSet},
	}
}
