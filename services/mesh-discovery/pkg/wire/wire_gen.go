// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"context"

	kubernetes_apps "github.com/solo-io/mesh-projects/pkg/clients/kubernetes/apps"
	kubernetes_core "github.com/solo-io/mesh-projects/pkg/clients/kubernetes/core"
	discovery_core "github.com/solo-io/mesh-projects/pkg/clients/zephyr/discovery"
	"github.com/solo-io/mesh-projects/pkg/common/docker"
	"github.com/solo-io/mesh-projects/services/common/multicluster/wire"
	mesh_workload "github.com/solo-io/mesh-projects/services/mesh-discovery/pkg/discovery/mesh-workload"
	"github.com/solo-io/mesh-projects/services/mesh-discovery/pkg/discovery/mesh/consul"
	"github.com/solo-io/mesh-projects/services/mesh-discovery/pkg/discovery/mesh/istio"
	"github.com/solo-io/mesh-projects/services/mesh-discovery/pkg/discovery/mesh/linkerd"
	"github.com/solo-io/mesh-projects/services/mesh-discovery/pkg/multicluster/controllers"
)

// Injectors from wire.go:

func InitializeDiscovery(ctx context.Context) (DiscoveryContext, error) {
	config, err := wire.LocalKubeConfigProvider()
	if err != nil {
		return DiscoveryContext{}, err
	}
	asyncManager, err := wire.LocalManagerProvider(ctx, config)
	if err != nil {
		return DiscoveryContext{}, err
	}
	asyncManagerController := wire.AsyncManagerControllerProvider(ctx, asyncManager)
	asyncManagerStartOptionsFunc := wire.LocalManagerStarterProvider(asyncManagerController)
	multiClusterDependencies := wire.MulticlusterDependenciesProvider(ctx, asyncManager, asyncManagerController, asyncManagerStartOptionsFunc)
	imageNameParser := docker.NewImageNameParser()
	istioMeshScanner := istio.NewIstioMeshScanner(imageNameParser)
	consulConnectInstallationScanner := consul.NewConsulConnectInstallationScanner(imageNameParser)
	consulConnectMeshScanner := consul.NewConsulMeshScanner(imageNameParser, consulConnectInstallationScanner)
	linkerdMeshScanner := linkerd.NewLinkerdMeshScanner(imageNameParser)
	replicaSetClientFactory := kubernetes_apps.ReplicaSetClientFactoryProvider()
	deploymentClientFactory := kubernetes_apps.DeploymentClientFactoryProvider()
	ownerFetcherFactory := mesh_workload.OwnerFetcherFactoryProvider()
	serviceClientFactory := kubernetes_core.ServiceClientFactoryProvider()
	meshServiceClientFactory := discovery_core.MeshServiceClientFactoryProvider()
	meshWorkloadClientFactory := discovery_core.MeshWorkloadClientFactoryProvider()
	podControllerFactory := controllers.NewPodControllerFactory()
	serviceControllerFactory := controllers.NewServiceControllerFactory()
	meshWorkloadControllerFactory := controllers.NewMeshWorkloadControllerFactory()
	deploymentControllerFactory := controllers.NewDeploymentControllerFactory()
	discoveryContext := DiscoveryContextProvider(multiClusterDependencies, istioMeshScanner, consulConnectMeshScanner, linkerdMeshScanner, replicaSetClientFactory, deploymentClientFactory, ownerFetcherFactory, serviceClientFactory, meshServiceClientFactory, meshWorkloadClientFactory, podControllerFactory, serviceControllerFactory, meshWorkloadControllerFactory, deploymentControllerFactory)
	return discoveryContext, nil
}
