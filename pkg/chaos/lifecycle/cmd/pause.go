package cmd

import (
	"context"
	"time"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// PauseParams holds the per-command parameters for the pause CLI subcommand.
type PauseParams struct {
	Duration time.Duration
	Limit    int
}

// NewPauseCLICommand initialize CLI pause command.
func NewPauseCLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func parsePauseParams(c cliflags.Flags, _ *chaos.GlobalParams) (PauseParams, error) {
	_ = "STUB: not implemented"
	return *new(PauseParams), nil
}

func buildPauseCommand(client container.Client, gp *chaos.GlobalParams, p PauseParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
