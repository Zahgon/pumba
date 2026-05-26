package cmd

import (
	"context"
	"time"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
	"github.com/urfave/cli"
)

// StressParams holds the per-command parameters for the stress CLI subcommand.
type StressParams struct {
	Image        string
	Pull         bool
	Stressors    string
	Duration     time.Duration
	Limit        int
	InjectCgroup bool
}

// NewStressCLICommand initialize CLI stress command.
func NewStressCLICommand(ctx context.Context, runtime chaos.Runtime) *cli.Command {
	_ = "STUB: not implemented"
	return nil
}

// :latest floats forward; --inject-cgroup needs /cg-inject which
// first shipped in 0.20.01. Pin to ":0.20.01" or newer if local
// cache predates that.

func parseStressParams(c cliflags.Flags, _ *chaos.GlobalParams) (StressParams, error) {
	_ = "STUB: not implemented"
	return *new(StressParams), nil
}

func buildStressCommand(client container.Client, gp *chaos.GlobalParams, p StressParams) (chaos.Command, error) {
	_ = "STUB: not implemented"
	return *new(chaos.Command), nil
}
