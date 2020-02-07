package auth

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	k8sapiv1 "k8s.io/api/core/v1"
	rbactypes "k8s.io/api/rbac/v1"
	kubeerrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Create a service account on a cluster that `targetClusterCfg` can reach
// Set up that service account with the indicated cluster roles
//go:generate mockgen -destination ./mocks/mock_remote_authority_manager.go github.com/solo-io/mesh-projects/pkg/auth RemoteAuthorityManager
type RemoteAuthorityManager interface {
	// creates a new service account in the cluster pointed to by the cfg at the name/namespace indicated by the ResourceRef,
	// and assigns it the given ClusterRoles
	// NB: if role assignment fails, the service account is left in the cluster; this is not an atomic operation
	ApplyRemoteServiceAccount(
		newServiceAccountRef *core.ResourceRef,
		roles []*rbactypes.ClusterRole) (*k8sapiv1.ServiceAccount, error)
}

func NewRemoteAuthorityManager(
	kubeClients kubernetes.Interface,
	rbacClient RbacClient, writeNamespace string) RemoteAuthorityManager {
	return &remoteAuthorityManager{
		saClient:   kubeClients.CoreV1().ServiceAccounts(writeNamespace),
		rbacClient: rbacClient,
	}
}

func NewRemoteAuthorityManagerForTest(saClient ServiceAccountClient, rbacClient RbacClient) RemoteAuthorityManager {
	return &remoteAuthorityManager{
		saClient:   saClient,
		rbacClient: rbacClient,
	}
}

type remoteAuthorityManager struct {
	saClient   ServiceAccountClient
	rbacClient RbacClient
}

func (r *remoteAuthorityManager) ApplyRemoteServiceAccount(
	newServiceAccountRef *core.ResourceRef,
	roles []*rbactypes.ClusterRole) (*k8sapiv1.ServiceAccount, error) {

	saToCreate := &k8sapiv1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: newServiceAccountRef.Name,
		},
	}
	newServiceAccount, err := r.saClient.Create(saToCreate)
	if err != nil {
		if kubeerrs.IsAlreadyExists(err) {
			newServiceAccount, err = r.saClient.Update(saToCreate)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	err = r.rbacClient.BindClusterRolesToServiceAccount(newServiceAccount, roles)
	if err != nil {
		return nil, err
	}

	return newServiceAccount, nil
}
