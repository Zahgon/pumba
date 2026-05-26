package containerd

import (
	"context"
	"time"

	ctr "github.com/alexei-led/pumba/pkg/container"
	containerd "github.com/containerd/containerd/v2/client"
)

const (
	sidecarCleanupTimeout = 30 * time.Second
	sidecarKillTimeout    = 5 * time.Second
)

// networkCapabilities are the Linux capabilities required for network manipulation.
var networkCapabilities = []string{"CAP_NET_ADMIN", "CAP_NET_RAW"}

// sidecarExec creates a short-lived sidecar container that shares the target
// container's network namespace and runs the given command+args inside it.
func (c *containerdClient) sidecarExec(ctx context.Context, target *ctr.Container, sidecarImage string, pull bool, command string, argsList [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Use context.WithoutCancel so cleanup succeeds even if the parent ctx is canceled.

// runSidecarCmd executes a single command inside a running sidecar task.
func (c *containerdClient) runSidecarCmd(ctx context.Context, task containerd.Task, command string, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// pullImage pulls an image via containerd.
func (c *containerdClient) pullImage(ctx context.Context, ref string) error {
	_ = "STUB: not implemented"
	return nil
}

// cleanupSidecar kills the task and removes the sidecar container and its snapshot.
func (c *containerdClient) cleanupSidecar(ctx context.Context, cntr containerd.Container) error {
	_ = "STUB: not implemented"
	return nil
}
