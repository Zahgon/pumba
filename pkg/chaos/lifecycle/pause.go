package lifecycle

import (
	"context"
	"time"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

// pauseClient is the narrow interface needed by the pause command.
type pauseClient interface {
	container.Lister
	PauseContainer(context.Context, *container.Container, bool) error
	UnpauseContainer(context.Context, *container.Container, bool) error
}

// `docker pause` command
type pauseCommand struct {
	client   pauseClient
	names    []string
	pattern  string
	labels   []string
	duration time.Duration
	limit    int
	dryRun   bool
}

// NewPauseCommand create new Pause Command instance
func NewPauseCommand(client pauseClient, params *chaos.GlobalParams, duration time.Duration, limit int) chaos.Command {
	_ = "STUB: not implemented"
	return *new(chaos.Command)
}

// Run pause command
func (p *pauseCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}

// if there are paused containers unpause them

// wait for specified duration and then unpause containers or unpause on ctx.Done()

// use context.WithoutCancel so cleanup succeeds even if the parent ctx is canceled

// unpause containers
func (p *pauseCommand) unpauseContainers(ctx context.Context, containers []*container.Container) error {
	_ = "STUB: not implemented"
	return nil
}
