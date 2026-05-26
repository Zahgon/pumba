package netem

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

// Parse rate
func parseRate(rate string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// `netem rate` command
type rateCommand struct {
	client         netemClient
	gp             *chaos.GlobalParams
	req            *container.NetemRequest
	limit          int
	rate           string
	packetOverhead int
	cellSize       int
	cellOverhead   int
}

// NewRateCommand create new netem rate command
func NewRateCommand(client netemClient,
	gp *chaos.GlobalParams,
	req *container.NetemRequest,
	limit int,
	rate string, // delay outgoing packets; in common units
	packetOverhead, // per packet overhead; in bytes
	cellSize, // cell size of the simulated link layer scheme
	cellOverhead int, // per cell overhead; in bytes
) (chaos.Command, error) {
	_ = "STUB: not implemented"
	// validate target egress rate
	return *new(chaos.Command), nil
}

// validate cell size

// Run netem rate command
func (n *rateCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *rateCommand) buildNetemCmd() []string { _ = "STUB: not implemented"; return nil }
