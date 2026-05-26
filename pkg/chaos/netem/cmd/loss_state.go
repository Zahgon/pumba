package cmd

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// LossStateParams holds the per-command parameters for the netem loss-state subcommand.
type LossStateParams struct {
	Base  *container.NetemRequest
	Limit int
	P13   float64
	P31   float64
	P32   float64
	P23   float64
	P14   float64
}

// NewLossStateCLICommand initialize CLI loss-state command.
func NewLossStateCLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

//nolint:mnd

//nolint:mnd

func parseLossStateParams(c cliflags.Flags, gp *chaos.GlobalParams) (LossStateParams, error) {
	_ = "STUB: not implemented"
	return *new(LossStateParams), nil
}

func buildLossStateCommand(client container.Client, gp *chaos.GlobalParams, p LossStateParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
