package k8s

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	k8stypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Perform a compile time check.
var _ Client = &KubeClient{}

type Client interface {
	PatchApply(context.Context, *unstructured.Unstructured) error
	GetStatefulSet(context.Context, string, string) (*appsv1.StatefulSet, error)
	Delete(context.Context, *unstructured.Unstructured) error
	GetSecret(context.Context, string, string) (*apiv1.Secret, error)
	DestinationRuleCRDExists(context.Context) (bool, error)
}

type KubeClient struct {
	client client.Client
	fieldManager string
}

func NewKubeClient(client client.Client, fieldManager string) Client {
	return &KubeClient{
		client:    client,
		fieldManager: fieldManager,
	}
}

func (c *KubeClient) PatchApply(ctx context.Context, object *unstructured.Unstructured) error {
	return c.client.Patch(ctx, object, client.Apply, &client.PatchOptions{
		Force:        pointer.Bool(true),
		FieldManager: c.fieldManager,
	})
}

func (c *KubeClient) Delete(ctx context.Context, object *unstructured.Unstructured) error {
	return client.IgnoreNotFound(c.client.Delete(ctx, object))
}

func (c *KubeClient) GetStatefulSet(ctx context.Context, name, namespace string) (*appsv1.StatefulSet, error) {
	nn := k8stypes.NamespacedName{
		Name: name,
		Namespace: namespace,
	}
	result := &appsv1.StatefulSet{}
	if err := c.client.Get(ctx, nn, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *KubeClient) GetSecret(ctx context.Context, name, namespace string) (*apiv1.Secret, error) {
	nn := k8stypes.NamespacedName{
		Name: name,
		Namespace: namespace,
	}
	result := &apiv1.Secret{}
	if err := c.client.Get(ctx, nn, result); err != nil {
		return nil, err
	}
	return nil, nil
}

func (c *KubeClient) DestinationRuleCRDExists(_ context.Context) (bool, error) {
	// @TODO: implement me
	return false, nil
}
