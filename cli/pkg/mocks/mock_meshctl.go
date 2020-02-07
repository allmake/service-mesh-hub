package cli_mocks

import (
	"bytes"
	"context"

	"github.com/golang/mock/gomock"
	"github.com/mattn/go-shellwords"
	"github.com/onsi/ginkgo"
	"github.com/solo-io/mesh-projects/cli/pkg/common"
	usage_mocks "github.com/solo-io/mesh-projects/cli/pkg/common/mocks"
	"k8s.io/client-go/rest"
)

// Build and execute the CLI app using the given clients
type MockMeshctl struct {
	// MUST be non-nil - we always need to produce a mocked master cluster verification client and a mocked usage reporter
	MockController *gomock.Controller

	// safe to leave nil if not needed
	MasterKubeConfig *rest.Config

	// safe to leave as nil if not needed
	Ctx context.Context

	Clients common.Clients

	KubeClients common.KubeClients
}

// call with the same string you would pass to the meshctl binary, i.e. "cluster register --service-account-name test123"
// returns the output produced by meshctl on stdout as a string
func (m MockMeshctl) Invoke(argString string) (stdout string, err error) {
	if m.MockController == nil {
		ginkgo.Fail("Must provide the ginkgo mock controller to mock meshctl")
	}

	buffer := &bytes.Buffer{}

	usageReporter := usage_mocks.NewMockClient(m.MockController)
	usageReporter.
		EXPECT().
		StartReportingUsage(m.Ctx, common.UsageReportingInterval).
		Return(nil)

	app := InitializeCLIWithMocks(
		m.Ctx,
		buffer,
		usageReporter,
		m.KubeClients.ClusterAuthorization,
		m.KubeClients.SecretWriter,
		m.Clients.ServerVersionClient,
		m.Clients.KubeLoader,
		m.Clients.MasterClusterVerifier,
		m.Clients.ReleaseAssetHelper)

	splitArgs, err := shellwords.Parse(argString)
	if err != nil {
		panic("Bad arg string: " + argString)
	}

	app.SetArgs(splitArgs)

	err = app.Execute()

	return buffer.String(), err
}
