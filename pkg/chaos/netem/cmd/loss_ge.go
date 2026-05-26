//nolint:dupl // Generic NewAction[P] enforces a uniform per-command shape; the residual similarity is intentional, not copy-paste.
package cmd

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// LossGEParams holds the per-command parameters for the netem loss-gemodel subcommand.
type LossGEParams struct {
	Base  *container.NetemRequest
	Limit int
	PG    float64
	PB    float64
	OneH  float64
	OneK  float64
}

// NewLossGECLICommand initialize CLI loss-gemodel command.
func NewLossGECLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

//nolint:mnd

//nolint:mnd

func parseLossGEParams(c cliflags.Flags, gp *chaos.GlobalParams) (LossGEParams, error) {
	_ = "STUB: not implemented"
	return *new(LossGEParams), nil
}

func buildLossGECommand(client container.Client, gp *chaos.GlobalParams, p LossGEParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
