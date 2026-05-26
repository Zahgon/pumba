package lifecycle

import (
	"context"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/container"
)

// execClient is the narrow interface needed by the exec command.
type execClient interface {
	container.Lister
	container.Executor
}

// `docker exec` command
type execCommand struct {
	client  execClient
	names   []string
	pattern string
	labels  []string
	command string
	args    []string
	limit   int
	dryRun  bool
}

// NewExecCommand create new Exec Command instance
func NewExecCommand(client execClient, params *chaos.GlobalParams, command string, args []string, limit int) chaos.Command {
	_ = "STUB: not implemented"
	return *new(chaos.Command)
}

// Run exec command
func (k *execCommand) Run(ctx context.Context, random bool) error {
	_ = "STUB: not implemented"
	return nil
}
