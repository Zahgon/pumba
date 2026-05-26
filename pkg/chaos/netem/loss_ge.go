package netem

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

// netem loss gemodel` (Gilbert-Elliot model) command
type lossGECommand struct {
	client netemClient
	gp     *chaos.GlobalParams
	req    *container.NetemRequest
	limit  int
	pg     float64
	pb     float64
	oneH   float64
	oneK   float64
}

// NewLossGECommand create new netem loss gemodel (Gilbert-Elliot) command
func NewLossGECommand(client netemClient,
	gp *chaos.GlobalParams,
	req *container.NetemRequest,
	limit int,
	pg, // Good State transition probability
	pb, // Bad State transition probability
	oneH, // loss probability in Bad state
	oneK float64, // loss probability in Good state
) (chaos.Command, error) {
	_ = "STUB: not implemented"
	// get pg - Good State transition probability
	return *new(chaos.Command), nil
}

// get pb - Bad State transition probability

// get (1-h) - loss probability in Bad state

// get (1-k) - loss probability in Good state

// Run netem loss state command
//
//nolint:dupl
func (n *lossGECommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *lossGECommand) buildNetemCmd() []string { _ = "STUB: not implemented"; return nil }
