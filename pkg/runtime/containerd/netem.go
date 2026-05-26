package containerd

import (
	"context"

	ctr "github.com/alexei-led/pumba/pkg/container"
)

// NetemContainer applies network emulation to a container by executing tc commands.
func (c *containerdClient) NetemContainer(ctx context.Context, req *ctr.NetemRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// StopNetemContainer removes network emulation from a container.
func (c *containerdClient) StopNetemContainer(ctx context.Context, req *ctr.NetemRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *containerdClient) runTCCommands(ctx context.Context, containerID string, commands [][]string) error {
	_ = "STUB: not implemented"
	return nil
}
