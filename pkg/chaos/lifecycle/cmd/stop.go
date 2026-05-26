package cmd

import (
	"context"
	"time"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// StopParams holds the per-command parameters for the stop CLI subcommand.
type StopParams struct {
	WaitTime int
	Limit    int
	Restart  bool
	Duration time.Duration
}

// NewStopCLICommand initialize CLI stop command.
func NewStopCLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func parseStopParams(c cliflags.Flags, _ *chaos.GlobalParams) (StopParams, error) {
	_ = "STUB: not implemented"
	return *new(StopParams), nil
}

func buildStopCommand(client container.Client, gp *chaos.GlobalParams, p StopParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
