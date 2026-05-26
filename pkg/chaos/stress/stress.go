package stress

import (
	"context"
	"time"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

// stressClient is the narrow interface needed by the stress command.
type stressClient interface {
	container.Lister
	container.Stressor
	StopContainerWithID(context.Context, string, time.Duration, bool) error
}

// `stress-ng` command
type stressCommand struct {
	client       stressClient
	names        []string
	pattern      string
	labels       []string
	image        string
	pull         bool
	stressors    []string
	duration     time.Duration
	limit        int
	injectCgroup bool
	dryRun       bool
}

const (
	defaultStopTimeout = 5 * time.Second
)

// NewStressCommand create new Kill stressCommand instance
func NewStressCommand(client stressClient, globalParams *chaos.GlobalParams, image string, pull bool, stressors string, duration time.Duration, limit int, injectCgroup bool) chaos.Command {
	_ = "STUB: not implemented"
	return *new(chaos.Command)
}

// Run stress command
func (s *stressCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *stressCommand) stressContainer(ctx context.Context, c *container.Container) error {
	_ = "STUB: not implemented"
	return nil
}

// cleanup must run even when parent ctx is canceled; preserve values but strip cancellation

// parent ctx may cancel simultaneously with the timer; strip cancellation for cleanup
