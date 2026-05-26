// Package containerd implements the container.Client interface for the containerd runtime.
package containerd

import (
	"context"
	"syscall"
	"time"

	ctr "github.com/alexei-led/pumba/pkg/container"
	containerd "github.com/containerd/containerd/v2/client"
)

const (
	defaultSocket    = "/run/containerd/containerd.sock"
	defaultNamespace = "k8s.io"
)

// containerdClient implements ctr.Client for the containerd runtime.
type containerdClient struct {
	client    apiClient
	namespace string
}

// NewClient creates a new containerd client connected to the given socket.
func NewClient(socket, namespace string) (ctr.Client, error) {
	_ = "STUB: not implemented"
	return *new(ctr.Client), nil
}

// Close releases the containerd client connection.
func (c *containerdClient) Close() error { _ = "STUB: not implemented"; return nil }

// nsCtx returns a context with the containerd namespace set.
func (c *containerdClient) nsCtx(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// resolveStopSignal parses the container's configured stop signal, falling back to SIGTERM.
func resolveStopSignal(container *ctr.Container) syscall.Signal {
	_ = "STUB: not implemented"
	return *new(syscall.Signal)
}

// forceKillTask kills and deletes a container's task for forced removal.
func (c *containerdClient) forceKillTask(ctx context.Context, cntr containerd.Container, id string) {
	_ = "STUB: not implemented"
	return
}

// ListContainers lists containers from containerd and applies the filter.
func (c *containerdClient) ListContainers(ctx context.Context, fn ctr.FilterFunc, opts ctr.ListOpts) ([]*ctr.Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StopContainer stops a container by sending its configured stop signal and waiting.
func (c *containerdClient) StopContainer(ctx context.Context, container *ctr.Container, timeout int, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// KillContainer sends a signal to the container's task.
func (c *containerdClient) KillContainer(ctx context.Context, container *ctr.Container, signal string, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// StartContainer starts a container's task.
func (c *containerdClient) StartContainer(ctx context.Context, container *ctr.Container, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// RestartContainer stops and starts a container's task.
func (c *containerdClient) RestartContainer(ctx context.Context, container *ctr.Container, timeout time.Duration, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveContainer deletes a container and optionally its task.
func (c *containerdClient) RemoveContainer(ctx context.Context, container *ctr.Container, opts ctr.RemoveOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// Try to delete with snapshot cleanup first; if that fails (e.g. Docker-managed
// containers in the moby namespace have no snapshot key), fall back to plain delete.
// Note: For Docker-managed containers, Docker daemon may react to the task being
// killed and clean up the container automatically, so "not found" is acceptable.

// PauseContainer pauses a container's task.
func (c *containerdClient) PauseContainer(ctx context.Context, container *ctr.Container, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpauseContainer resumes a paused container's task.
func (c *containerdClient) UnpauseContainer(ctx context.Context, container *ctr.Container, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// StopContainerWithID stops a container by ID.
func (c *containerdClient) StopContainerWithID(ctx context.Context, containerID string, timeout time.Duration, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// ExecContainer executes a command inside a running container.
func (c *containerdClient) ExecContainer(ctx context.Context, container *ctr.Container, command string, args []string, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}
