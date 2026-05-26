package cmd

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// ExecParams holds the per-command parameters for the exec CLI subcommand.
type ExecParams struct {
	Command string
	Args    []string
	Limit   int
}

// NewExecCLICommand initialize CLI exec command.
func NewExecCLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func parseExecParams(c cliflags.Flags, _ *chaos.GlobalParams) (ExecParams, error) {
	_ = "STUB: not implemented"
	return *new(ExecParams), nil
}

func buildExecCommand(client container.Client, gp *chaos.GlobalParams, p ExecParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
