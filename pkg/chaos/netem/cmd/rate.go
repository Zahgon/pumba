//nolint:dupl // Generic NewAction[P] enforces a uniform per-command shape; the residual similarity is intentional, not copy-paste.
package cmd

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// RateParams holds the per-command parameters for the netem rate subcommand.
type RateParams struct {
	Base           *container.NetemRequest
	Limit          int
	Rate           string
	PacketOverhead int
	CellSize       int
	CellOverhead   int
}

// NewRateCLICommand initialize CLI rate command.
func NewRateCLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func parseRateParams(c cliflags.Flags, gp *chaos.GlobalParams) (RateParams, error) {
	_ = "STUB: not implemented"
	return *new(RateParams), nil
}

func buildRateCommand(client container.Client, gp *chaos.GlobalParams, p RateParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
