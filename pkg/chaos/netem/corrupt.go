package netem

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

// `netem corrupt` command
type corruptCommand struct {
	client      netemClient
	gp          *chaos.GlobalParams
	req         *container.NetemRequest
	limit       int
	percent     float64
	correlation float64
}

// NewCorruptCommand create new netem corrupt command
func NewCorruptCommand(client netemClient,
	gp *chaos.GlobalParams,
	req *container.NetemRequest,
	limit int,
	percent, // corrupt percent
	correlation float64, // corrupt correlation
) (chaos.Command, error) {
	_ = "STUB: not implemented"
	// get netem corrupt percent
	return *new(chaos.Command), nil
}

// get netem corrupt variation

// Run netem corrupt command
//
//nolint:dupl
func (n *corruptCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *corruptCommand) buildNetemCmd() []string { _ = "STUB: not implemented"; return nil }
