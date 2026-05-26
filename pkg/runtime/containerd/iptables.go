package containerd

import (
	"context"

	ctr "github.com/alexei-led/pumba/pkg/container"
)

// IPTablesContainer applies iptables rules to a container.
//
//nolint:dupl // intentionally parallel to StopIPTablesContainer; install/remove use identical IPTables commands on this runtime
func (c *containerdClient) IPTablesContainer(ctx context.Context, req *ctr.IPTablesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// StopIPTablesContainer removes iptables rules from a container.
//
//nolint:dupl // intentionally parallel to IPTablesContainer; install/remove use identical IPTables commands on this runtime
func (c *containerdClient) StopIPTablesContainer(ctx context.Context, req *ctr.IPTablesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *containerdClient) runIPTablesCommands(ctx context.Context, containerID string, commands [][]string) error {
	_ = "STUB: not implemented"
	return nil
}
