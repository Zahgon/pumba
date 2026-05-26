//nolint:dupl // Generic NewAction[P] enforces a uniform per-command shape; the residual similarity is intentional, not copy-paste.
package cmd

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// CorruptParams holds the per-command parameters for the netem corrupt subcommand.
type CorruptParams struct {
	Base        *container.NetemRequest
	Limit       int
	Percent     float64
	Correlation float64
}

// NewCorruptCLICommand initialize CLI corrupt command.
func NewCorruptCLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func parseCorruptParams(c cliflags.Flags, gp *chaos.GlobalParams) (CorruptParams, error) {
	_ = "STUB: not implemented"
	return *new(CorruptParams), nil
}

func buildCorruptCommand(client container.Client, gp *chaos.GlobalParams, p CorruptParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
