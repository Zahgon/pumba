package chaos

import (
	"context"
	"time"

	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
)

const (
	// Re2Prefix re2 regexp string prefix
	Re2Prefix = "re2:"
)

// Runtime returns the container client to use for chaos execution. Builders
// receive a Runtime factory rather than a client value so that client
// construction can be deferred until after global flag parsing while still
// keeping the dependency visible in every constructor signature.
type Runtime func() container.Client

// Command chaos command
type Command interface {
	Run(ctx context.Context, random bool) error
}

// GlobalParams global parameters passed through CLI flags
type GlobalParams struct {
	Random     bool
	Labels     []string
	Pattern    string
	Names      []string
	Interval   time.Duration
	DryRun     bool
	SkipErrors bool
}

// splitLabels splits comma-separated label values into individual labels.
// This supports both "--label k1=v1 --label k2=v2" and "--label k1=v1,k2=v2" syntax.
func splitLabels(raw []string) []string { _ = "STUB: not implemented"; return nil }

// ParseGlobalParams parses application-level flags from any cliflags.Flags
// implementation. Reads global flags via c.Global() so the caller may pass the
// subcommand-level Flags directly.
func ParseGlobalParams(c cliflags.Flags) *GlobalParams { _ = "STUB: not implemented"; return nil }

// get names list of filter pattern from command line
func getNamesOrPattern(c cliflags.Flags) ([]string, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// no Args means ALL containers

// more than one argument, assume that this a list of names

// RunChaosCommand run chaos command in go routine
func RunChaosCommand(topContext context.Context, command Command, params *GlobalParams) error {
	_ = "STUB: not implemented"
	// create Time channel for specified interval
	return nil
}

// handle the 'chaos' command

// cancel current context on exit

// run chaos command

// run chaos function

// wait for next timer tick or cancel

// not to leak the goroutine

// not to leak the goroutine
