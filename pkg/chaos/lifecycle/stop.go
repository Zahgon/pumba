package lifecycle

import (
	"context"
	"time"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

const (
	// DeafultWaitTime time to wait before stopping container (in seconds)
	DeafultWaitTime = 5
)

// stopClient is the narrow interface needed by the stop command.
type stopClient interface {
	container.Lister
	StopContainer(context.Context, *container.Container, int, bool) error
	StartContainer(context.Context, *container.Container, bool) error
}

// `docker stop` command
type stopCommand struct {
	client   stopClient
	names    []string
	pattern  string
	labels   []string
	restart  bool
	duration time.Duration
	waitTime int
	limit    int
	dryRun   bool
}

// NewStopCommand create new Stop Command instance
func NewStopCommand(client stopClient, params *chaos.GlobalParams, restart bool, duration time.Duration, waitTime, limit int) chaos.Command {
	_ = "STUB: not implemented"
	return *new(chaos.Command)
}

// Run stop command
func (s *stopCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}

// if there are stopped containers and want to (re)start ...

// wait for specified duration and then start containers or start on ctx.Done()

// use context.WithoutCancel so cleanup succeeds even if the parent ctx is canceled

// start previously stopped containers after duration on exit
func (s *stopCommand) startStoppedContainers(ctx context.Context, containers []*container.Container) error {
	_ = "STUB: not implemented"
	return nil
}
