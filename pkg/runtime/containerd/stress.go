package containerd

import (
	"context"
	"time"

	ctr "github.com/alexei-led/pumba/pkg/container"
)

// StressContainer runs stress-ng to stress a container.
// Mode selection:
//   - Sidecar.Image == "": direct exec inside the target container
//   - Sidecar.Image != "" && !InjectCgroup: sidecar with /stress-ng in target's cgroup parent
//   - Sidecar.Image != "" && InjectCgroup: sidecar with /cg-inject injecting into target's cgroup
func (c *containerdClient) StressContainer(ctx context.Context, req *ctr.StressRequest) (*ctr.StressResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// stressDirectExec runs stress-ng directly inside the target container via exec.
func (c *containerdClient) stressDirectExec(ctx context.Context, container *ctr.Container,
	stressors []string, duration time.Duration) (string, <-chan string, <-chan error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// stressSidecar creates a long-lived sidecar container running stress-ng (or cg-inject)
// as its main process. Returns the sidecar ID and output/error channels. A goroutine waits
// for the task to exit and performs full cleanup (task delete + container/snapshot removal).
func (c *containerdClient) stressSidecar(
	ctx context.Context,
	target *ctr.Container,
	sidecarImage string,
	stressors []string,
	injectCgroup bool,
	pull bool,
) (string, <-chan string, <-chan error, error) {
	_ = "STUB: not implemented"
	return "", nil, nil, nil
}
