package k8s

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	ERRNamespaceConflict = fmt.Errorf("namespace already exists")
	ERRNamespaceNotExist = fmt.Errorf("namespace not exist")
)

type NamespaceManager struct {
	ClientSet *kubernetes.Clientset
}

func (m *NamespaceManager) CreateNamespace(ctx context.Context, name string) error {
	namespaceClient := m.ClientSet.CoreV1().Namespaces()
	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	_, err := namespaceClient.Create(ctx, namespace, metav1.CreateOptions{})
	if errors.IsAlreadyExists(err) {
		return ERRNamespaceConflict
	}

	return err
}

func (m *NamespaceManager) GetNamespace(ctx context.Context, name string) (*corev1.Namespace, error) {
	namespaceClient := m.ClientSet.CoreV1().Namespaces()
	namespace, err := namespaceClient.Get(ctx, name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return nil, ERRNamespaceNotExist
	}

	return namespace, err
}
