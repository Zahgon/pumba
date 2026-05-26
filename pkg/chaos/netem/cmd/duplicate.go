//nolint:dupl // Generic NewAction[P] enforces a uniform per-command shape; the residual similarity is intentional, not copy-paste.
package cmd

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// DuplicateParams holds the per-command parameters for the netem duplicate subcommand.
type DuplicateParams struct {
	Base        *container.NetemRequest
	Limit       int
	Percent     float64
	Correlation float64
}

// NewDuplicateCLICommand initialize CLI duplicate command.
func NewDuplicateCLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

func parseDuplicateParams(c cliflags.Flags, gp *chaos.GlobalParams) (DuplicateParams, error) {
	_ = "STUB: not implemented"
	return *new(DuplicateParams), nil
}

func buildDuplicateCommand(client container.Client, gp *chaos.GlobalParams, p DuplicateParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
