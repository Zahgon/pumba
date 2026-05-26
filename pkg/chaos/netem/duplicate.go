package netem

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

const (
	duplicateCmd = "duplicate"
)

// `netem duplicate` command
type duplicateCommand struct {
	client      netemClient
	gp          *chaos.GlobalParams
	req         *container.NetemRequest
	limit       int
	percent     float64
	correlation float64
}

// NewDuplicateCommand create new netem duplicate command
func NewDuplicateCommand(client netemClient,
	gp *chaos.GlobalParams,
	req *container.NetemRequest,
	limit int,
	percent, // duplicate percent
	correlation float64, // duplicate correlation
) (chaos.Command, error) {
	_ = "STUB: not implemented"
	// get netem duplicate percent
	return *new(chaos.Command), nil
}

// get netem duplicate variation

// Run netem duplicate command
//
//nolint:dupl
func (n *duplicateCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *duplicateCommand) buildNetemCmd() []string { _ = "STUB: not implemented"; return nil }
