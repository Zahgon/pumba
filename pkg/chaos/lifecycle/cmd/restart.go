package cmd

import (
	"context"
	"time"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// RestartParams holds the per-command parameters for the restart CLI subcommand.
type RestartParams struct {
	Timeout time.Duration
	Limit   int
}

// NewRestartCLICommand initialize CLI restart command.
func NewRestartCLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func parseRestartParams(c cliflags.Flags, _ *chaos.GlobalParams) (RestartParams, error) {
	_ = "STUB: not implemented"
	return *new(RestartParams), nil
}

func buildRestartCommand(client container.Client, gp *chaos.GlobalParams, p RestartParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
