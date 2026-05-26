package docker

import (
	"context"

	ctr "github.com/alexei-led/pumba/pkg/container"
)

// IPTablesContainer injects sidecar iptables container into the given container network namespace
func (client dockerClient) IPTablesContainer(ctx context.Context, req *ctr.IPTablesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// StopIPTablesContainer stops the iptables container injected into the given container network namespace
func (client dockerClient) StopIPTablesContainer(ctx context.Context, req *ctr.IPTablesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (client dockerClient) ipTablesContainer(ctx context.Context, req *ctr.IPTablesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (client dockerClient) ipTablesContainerWithIPFilter(ctx context.Context, req *ctr.IPTablesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// use docker client ExecStart to run iptables rules to filter network

// See more about the iptables statistics extension: https://www.man7.org/linux/man-pages/man8/iptables-extensions.8.html
// # drop traffic to a specific source address

// # drop traffic to a specific destination address

// # drop traffic to a specific source port

// # drop traffic to a specific destination port

func (client dockerClient) ipTablesCommands(ctx context.Context, c *ctr.Container, argsList [][]string, tcimg string, pull bool) error {
	_ = "STUB: not implemented"
	return nil
}
