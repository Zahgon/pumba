package lifecycle

import (
	"context"
	"time"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

// restartClient is the narrow interface needed by the restart command.
type restartClient interface {
	container.Lister
	RestartContainer(context.Context, *container.Container, time.Duration, bool) error
}

// `docker restart` command
type restartCommand struct {
	client  restartClient
	names   []string
	pattern string
	labels  []string
	timeout time.Duration
	limit   int
	dryRun  bool
}

// NewRestartCommand create new Restart Command instance
func NewRestartCommand(client restartClient, params *chaos.GlobalParams, timeout time.Duration, limit int) chaos.Command {
	_ = "STUB: not implemented"
	return *new(chaos.Command)
}

// Run restart command
func (k *restartCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}
