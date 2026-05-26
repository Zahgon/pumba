//nolint:dupl // Generic NewAction[P] enforces a uniform per-command shape; the residual similarity is intentional, not copy-paste.
package cmd

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// DelayParams holds the per-command parameters for the netem delay subcommand.
type DelayParams struct {
	Base         *container.NetemRequest
	Limit        int
	Time         int
	Jitter       int
	Correlation  float64
	Distribution string
}

// NewDelayCLICommand initialize CLI delay command.
func NewDelayCLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

//nolint:mnd

//nolint:mnd

//nolint:mnd

func parseDelayParams(c cliflags.Flags, gp *chaos.GlobalParams) (DelayParams, error) {
	_ = "STUB: not implemented"
	return *new(DelayParams), nil
}

func buildDelayCommand(client container.Client, gp *chaos.GlobalParams, p DelayParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
