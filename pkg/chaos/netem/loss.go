package netem

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

// `netem loss` command
type lossCommand struct {
	client      netemClient
	gp          *chaos.GlobalParams
	req         *container.NetemRequest
	limit       int
	percent     float64
	correlation float64
}

// NewLossCommand create new netem loss command
func NewLossCommand(client netemClient,
	gp *chaos.GlobalParams,
	req *container.NetemRequest,
	limit int,
	percent, // loss percent
	correlation float64, // loss correlation
) (chaos.Command, error) {
	_ = "STUB: not implemented"
	// get netem loss percent
	return *new(chaos.Command), nil
}

// get netem loss variation

// Run netem loss command
func (n *lossCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *lossCommand) buildNetemCmd() []string { _ = "STUB: not implemented"; return nil }
