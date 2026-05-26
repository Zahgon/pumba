package netem

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

// `netem loss state` command
type lossStateCommand struct {
	client netemClient
	gp     *chaos.GlobalParams
	req    *container.NetemRequest
	limit  int
	p13    float64
	p31    float64
	p32    float64
	p23    float64
	p14    float64
}

// NewLossStateCommand create new netem loss state command
func NewLossStateCommand(client netemClient,
	gp *chaos.GlobalParams,
	req *container.NetemRequest,
	limit int,
	p13, // probability to go from state (1) to state (3)
	p31, // probability to go from state (3) to state (1)
	p32, // probability to go from state (3) to state (2)
	p23, // probability to go from state (2) to state (3)
	p14 float64, // probability to go from state (1) to state (4)
) (chaos.Command, error) {
	_ = "STUB: not implemented"
	// validate p13
	return *new(chaos.Command), nil
}

// validate p31

// validate p32

// vaidate p23

// validate p14

// Run netem loss state command
//
//nolint:dupl
func (n *lossStateCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *lossStateCommand) buildNetemCmd() []string { _ = "STUB: not implemented"; return nil }
