package cleanup

import (
	"context"

	"github.com/solo-io/service-mesh-hub/cli/pkg/common/exec"
)

func DemoCleanup(ctx context.Context, runner exec.Runner) error {
	return runner.Run("bash", "-c", cleanupDemoScript)
}

const (
	cleanupDemoScript = `
kind get clusters | while read -r r; do kind delete cluster --name "$r"; done
exit 0
`
)
