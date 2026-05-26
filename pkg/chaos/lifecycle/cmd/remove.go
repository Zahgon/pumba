package cmd

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// RemoveParams holds the per-command parameters for the rm CLI subcommand.
type RemoveParams struct {
	Force   bool
	Links   bool
	Volumes bool
	Limit   int
}

// NewRemoveCLICommand initialize CLI remove command.
func NewRemoveCLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func parseRemoveParams(c cliflags.Flags, _ *chaos.GlobalParams) (RemoveParams, error) {
	_ = "STUB: not implemented"
	return *new(RemoveParams), nil
}

func buildRemoveCommand(client container.Client, gp *chaos.GlobalParams, p RemoveParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
