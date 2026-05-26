package containerd

import (
	"context"

	ctr "github.com/alexei-led/pumba/pkg/container"
	containerd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/pkg/oci"
)

// buildStressSpecOpts builds OCI spec options for the stress sidecar container.
// For inject-cgroup mode (uses cgroupPath): runs /cg-inject with host cgroupns and /sys/fs/cgroup bind mount.
// For default sidecar mode (uses cgroupParent + sidecarID): runs /stress-ng directly, placed in a child cgroup
// under the target's cgroup parent. The child path format depends on the cgroup driver (systemd vs cgroupfs).
func buildStressSpecOpts(image containerd.Image, stressors []string, cgroupPath, cgroupParent, sidecarID string, injectCgroup bool) []oci.SpecOpts {
	_ = "STUB: not implemented"
	return nil
}

// createStressSidecar resolves the target cgroup path, optionally pulls the image, and creates and starts the stress sidecar container.
func (c *containerdClient) createStressSidecar(
	ctx context.Context,
	target *ctr.Container,
	sidecarImage string,
	stressors []string,
	injectCgroup bool,
	pull bool,
) (string, containerd.Container, containerd.Task, <-chan containerd.ExitStatus, error) {
	_ = "STUB: not implemented"
	return "", *new(containerd.Container), *new(containerd.Task), nil, nil
}

// startSidecarTask creates, registers wait, and starts a task on the given container.
func (c *containerdClient) startSidecarTask(ctx context.Context, cntr containerd.Container) (containerd.Task, <-chan containerd.ExitStatus, error) {
	_ = "STUB: not implemented"
	return *new(containerd.Task), nil, nil
}

// waitStressSidecar waits for the stress task to exit, performs cleanup, and reports the result.
func (c *containerdClient) waitStressSidecar(
	ctx context.Context,
	sidecarID string,
	sidecarContainer containerd.Container,
	task containerd.Task,
	waitCh <-chan containerd.ExitStatus,
	outCh chan<- string,
	errCh chan<- error,
) {
	_ = "STUB: not implemented"
	return
}

// deleteContainer deletes cntr and its snapshot.
// The caller must provide a context with a timeout and the containerd namespace already set.
func (c *containerdClient) deleteContainer(ctx context.Context, cntr containerd.Container) {
	_ = "STUB: not implemented"
	return
}
