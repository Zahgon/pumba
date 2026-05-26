package docker

import (
	"context"
	"time"

	ctr "github.com/alexei-led/pumba/pkg/container"
)

const (
	defaultStopSignal = "SIGTERM"
	defaultKillSignal = "SIGKILL"
)

// KillContainer kills a container with the given signal
func (client dockerClient) KillContainer(ctx context.Context, c *ctr.Container, signal string, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// RestartContainer restarts a container
func (client dockerClient) RestartContainer(ctx context.Context, c *ctr.Container, timeout time.Duration, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// StopContainer stops a container
func (client dockerClient) StopContainer(ctx context.Context, c *ctr.Container, timeout int, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for container to exit, but proceed anyway after the timeout elapses

// failed to stop gracefully - going to kill target container

// Wait for container to be removed

// StopContainerWithID stops a container with a timeout
func (client dockerClient) StopContainerWithID(ctx context.Context, containerID string, timeout time.Duration, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// StartContainer starts a container
func (client dockerClient) StartContainer(ctx context.Context, c *ctr.Container, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveContainer removes a container
func (client dockerClient) RemoveContainer(ctx context.Context, c *ctr.Container, opts ctr.RemoveOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// PauseContainer pauses a container main process
func (client dockerClient) PauseContainer(ctx context.Context, c *ctr.Container, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// UnpauseContainer unpauses a container main process
func (client dockerClient) UnpauseContainer(ctx context.Context, c *ctr.Container, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (client dockerClient) waitForStop(ctx context.Context, c *ctr.Container, waitTime int) error {
	_ = "STUB: not implemented"
	// check status every 100 ms
	return nil
}

// timeout after waitTime seconds
