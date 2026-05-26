package netem

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

var (
	// DelayDistribution netem delay distributions
	delayDistribution = []string{"", "uniform", "normal", "pareto", "paretonormal"}
)

// `netem delay` command
type delayCommand struct {
	client       netemClient
	gp           *chaos.GlobalParams
	req          *container.NetemRequest
	limit        int
	time         int
	jitter       int
	correlation  float64
	distribution string
}

// NewDelayCommand create new netem delay command
func NewDelayCommand(client netemClient,
	gp *chaos.GlobalParams,
	req *container.NetemRequest,
	limit int,
	delay, // delay time
	jitter int, // delay jitter
	correlation float64, // delay correlation
	distribution string, // delay distribution
) (chaos.Command, error) {
	_ = "STUB: not implemented"
	// check delay time
	return *new(chaos.Command), nil
}

// get delay variation

// get delay variation

// get distribution

// Run netem delay command
func (n *delayCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *delayCommand) buildNetemCmd() []string { _ = "STUB: not implemented"; return nil }
