package v1

import (
	"context"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	. "k8s.io/api/apps/v1"
)

// clienset for the apps/v1 APIs
type Clientset interface {
	// clienset for the apps/v1/v1 APIs
	Deployments() DeploymentClient
	// clienset for the apps/v1/v1 APIs
	ReplicaSets() ReplicaSetClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (*clientSet, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) *clientSet {
	return &clientSet{client: client}
}

// clienset for the apps/v1/v1 APIs
func (c *clientSet) Deployments() DeploymentClient {
	return NewDeploymentClient(c.client)
}

// clienset for the apps/v1/v1 APIs
func (c *clientSet) ReplicaSets() ReplicaSetClient {
	return NewReplicaSetClient(c.client)
}

// Reader knows how to read and list Deployments.
type DeploymentReader interface {
	// Get retrieves a Deployment for the given object key
	GetDeployment(ctx context.Context, key client.ObjectKey) (*Deployment, error)

	// List retrieves list of Deployments for a given namespace and list options.
	ListDeployment(ctx context.Context, opts ...client.ListOption) (*DeploymentList, error)
}

// Writer knows how to create, delete, and update Deployments.
type DeploymentWriter interface {
	// Create saves the Deployment object.
	CreateDeployment(ctx context.Context, obj *Deployment, opts ...client.CreateOption) error

	// Delete deletes the Deployment object.
	DeleteDeployment(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given Deployment object.
	UpdateDeployment(ctx context.Context, obj *Deployment, opts ...client.UpdateOption) error

	// Patch patches the given Deployment object.
	PatchDeployment(ctx context.Context, obj *Deployment, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all Deployment objects matching the given options.
	DeleteAllOfDeployment(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a Deployment object.
type DeploymentStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given Deployment object.
	UpdateDeploymentStatus(ctx context.Context, obj *Deployment, opts ...client.UpdateOption) error

	// Patch patches the given Deployment object's subresource.
	PatchDeploymentStatus(ctx context.Context, obj *Deployment, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on Deployments.
type DeploymentClient interface {
	DeploymentReader
	DeploymentWriter
	DeploymentStatusWriter
}

type deploymentClient struct {
	client client.Client
}

func NewDeploymentClient(client client.Client) *deploymentClient {
	return &deploymentClient{client: client}
}

func (c *deploymentClient) GetDeployment(ctx context.Context, key client.ObjectKey) (*Deployment, error) {
	obj := &Deployment{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *deploymentClient) ListDeployment(ctx context.Context, opts ...client.ListOption) (*DeploymentList, error) {
	list := &DeploymentList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *deploymentClient) CreateDeployment(ctx context.Context, obj *Deployment, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *deploymentClient) DeleteDeployment(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &Deployment{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *deploymentClient) UpdateDeployment(ctx context.Context, obj *Deployment, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *deploymentClient) PatchDeployment(ctx context.Context, obj *Deployment, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *deploymentClient) DeleteAllOfDeployment(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &Deployment{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *deploymentClient) UpdateDeploymentStatus(ctx context.Context, obj *Deployment, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *deploymentClient) PatchDeploymentStatus(ctx context.Context, obj *Deployment, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list ReplicaSets.
type ReplicaSetReader interface {
	// Get retrieves a ReplicaSet for the given object key
	GetReplicaSet(ctx context.Context, key client.ObjectKey) (*ReplicaSet, error)

	// List retrieves list of ReplicaSets for a given namespace and list options.
	ListReplicaSet(ctx context.Context, opts ...client.ListOption) (*ReplicaSetList, error)
}

// Writer knows how to create, delete, and update ReplicaSets.
type ReplicaSetWriter interface {
	// Create saves the ReplicaSet object.
	CreateReplicaSet(ctx context.Context, obj *ReplicaSet, opts ...client.CreateOption) error

	// Delete deletes the ReplicaSet object.
	DeleteReplicaSet(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given ReplicaSet object.
	UpdateReplicaSet(ctx context.Context, obj *ReplicaSet, opts ...client.UpdateOption) error

	// Patch patches the given ReplicaSet object.
	PatchReplicaSet(ctx context.Context, obj *ReplicaSet, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all ReplicaSet objects matching the given options.
	DeleteAllOfReplicaSet(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a ReplicaSet object.
type ReplicaSetStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given ReplicaSet object.
	UpdateReplicaSetStatus(ctx context.Context, obj *ReplicaSet, opts ...client.UpdateOption) error

	// Patch patches the given ReplicaSet object's subresource.
	PatchReplicaSetStatus(ctx context.Context, obj *ReplicaSet, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on ReplicaSets.
type ReplicaSetClient interface {
	ReplicaSetReader
	ReplicaSetWriter
	ReplicaSetStatusWriter
}

type replicaSetClient struct {
	client client.Client
}

func NewReplicaSetClient(client client.Client) *replicaSetClient {
	return &replicaSetClient{client: client}
}

func (c *replicaSetClient) GetReplicaSet(ctx context.Context, key client.ObjectKey) (*ReplicaSet, error) {
	obj := &ReplicaSet{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *replicaSetClient) ListReplicaSet(ctx context.Context, opts ...client.ListOption) (*ReplicaSetList, error) {
	list := &ReplicaSetList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *replicaSetClient) CreateReplicaSet(ctx context.Context, obj *ReplicaSet, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *replicaSetClient) DeleteReplicaSet(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &ReplicaSet{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *replicaSetClient) UpdateReplicaSet(ctx context.Context, obj *ReplicaSet, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *replicaSetClient) PatchReplicaSet(ctx context.Context, obj *ReplicaSet, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *replicaSetClient) DeleteAllOfReplicaSet(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &ReplicaSet{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *replicaSetClient) UpdateReplicaSetStatus(ctx context.Context, obj *ReplicaSet, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *replicaSetClient) PatchReplicaSetStatus(ctx context.Context, obj *ReplicaSet, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}
