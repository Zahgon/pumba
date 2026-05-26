package containerd

import (
	"context"
	"sync/atomic"
	"syscall"
	"time"

	containerd "github.com/containerd/containerd/v2/client"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

// execCounter generates unique IDs for exec and sidecar operations.
var execCounter atomic.Uint64

// signalMap maps signal names to syscall signals.
var signalMap = map[string]syscall.Signal{
	"SIGABRT": syscall.SIGABRT,
	"SIGALRM": syscall.SIGALRM,
	"SIGCONT": syscall.SIGCONT,
	"SIGHUP":  syscall.SIGHUP,
	"SIGINT":  syscall.SIGINT,
	"SIGKILL": syscall.SIGKILL,
	"SIGPIPE": syscall.SIGPIPE,
	"SIGQUIT": syscall.SIGQUIT,
	"SIGSTOP": syscall.SIGSTOP,
	"SIGTERM": syscall.SIGTERM,
	"SIGTRAP": syscall.SIGTRAP,
	"SIGUSR1": syscall.SIGUSR1,
	"SIGUSR2": syscall.SIGUSR2,
}

func parseSignal(signal string) (syscall.Signal, error) {
	_ = "STUB: not implemented"
	return *new(syscall.Signal), nil
}

// killTimeout is the maximum time to wait for SIGKILL to take effect.
const killTimeout = 30 * time.Second

func (c *containerdClient) getTask(ctx context.Context, containerID string) (containerd.Task, error) {
	_ = "STUB: not implemented"
	return *new(containerd.Task), nil
}

func (c *containerdClient) stopTask(ctx context.Context, containerID string, signal syscall.Signal, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *containerdClient) killTask(ctx context.Context, containerID, signal string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *containerdClient) startTask(ctx context.Context, containerID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *containerdClient) pauseTask(ctx context.Context, containerID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *containerdClient) resumeTask(ctx context.Context, containerID string) error {
	_ = "STUB: not implemented"
	return nil
}

// execTask runs a command inside a containerd task and waits for completion.
// Handles the full exec lifecycle: create, wait, start, collect exit status, delete.
func execTask(ctx context.Context, task containerd.Task, pspec *specs.Process, execID, description string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *containerdClient) execInContainer(ctx context.Context, containerID, command string, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
