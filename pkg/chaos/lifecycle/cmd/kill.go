package cmd

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// KillParams holds the per-command parameters for the kill CLI subcommand.
type KillParams struct {
	Signal string
	Limit  int
}

// NewKillCLICommand initialize CLI kill command.
func NewKillCLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func parseKillParams(c cliflags.Flags, _ *chaos.GlobalParams) (KillParams, error) {
	_ = "STUB: not implemented"
	return *new(KillParams), nil
}

func buildKillCommand(client container.Client, gp *chaos.GlobalParams, p KillParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
