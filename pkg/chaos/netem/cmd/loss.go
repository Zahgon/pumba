//nolint:dupl // Generic NewAction[P] enforces a uniform per-command shape; the residual similarity is intentional, not copy-paste.
package cmd

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// LossParams holds the per-command parameters for the netem loss subcommand.
type LossParams struct {
	Base        *container.NetemRequest
	Limit       int
	Percent     float64
	Correlation float64
}

// NewLossCLICommand initialize CLI loss command.
func NewLossCLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func parseLossParams(c cliflags.Flags, gp *chaos.GlobalParams) (LossParams, error) {
	_ = "STUB: not implemented"
	return *new(LossParams), nil
}

func buildLossCommand(client container.Client, gp *chaos.GlobalParams, p LossParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
