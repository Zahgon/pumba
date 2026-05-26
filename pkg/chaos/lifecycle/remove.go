package lifecycle

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

// removeClient is the narrow interface needed by the remove command.
type removeClient interface {
	container.Lister
	RemoveContainer(context.Context, *container.Container, container.RemoveOpts) error
}

// `docker rm` command
type removeCommand struct {
	client  removeClient
	names   []string
	pattern string
	labels  []string
	opts    container.RemoveOpts
	limit   int
}

// NewRemoveCommand create new Kill Command instance
func NewRemoveCommand(client removeClient, params *chaos.GlobalParams, force, links, volumes bool, limit int) chaos.Command {
	_ = "STUB: not implemented"
	return *new(chaos.Command)
}

// Run remove command
func (r *removeCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}
