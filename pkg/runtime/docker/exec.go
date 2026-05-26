package docker

import (
	"context"

	ctr "github.com/alexei-led/pumba/pkg/container"
)

// ExecContainer executes a command in a container
func (client dockerClient) ExecContainer(ctx context.Context, c *ctr.Container, command string, args []string, dryrun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// runExecAttached starts a pre-created exec by attaching to it and draining
// stdout/stderr until the exec completes. Podman's Docker-compat API rejects
// ContainerExecStart with empty ExecStartOptions ("must provide at least one
// stream to attach to"); Docker accepts it. ContainerExecAttach works on both.
func (client dockerClient) runExecAttached(ctx context.Context, execID string) error {
	_ = "STUB: not implemented"
	return nil
}

// execute command on container
func (client dockerClient) execOnContainer(ctx context.Context, c *ctr.Container, execCmd string, execArgs []string, privileged bool) error {
	_ = "STUB: not implemented"
	return nil
}

// trim all spaces from cmd

// check if command exists inside target container

// if command found execute it

// prepare exec config

// execute the command
