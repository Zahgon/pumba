package netem

import (
	"context"
	"time"

	"github.com/alexei-led/pumba/pkg/container"
)

// netemClient is the narrow interface needed by all netem commands.
type netemClient interface {
	container.Lister
	container.Netem
}

// cleanupTimeout caps how long the netem-cleanup sidecar cycle is allowed
// to run after abort or scheduled stop. Independent of --duration so a
// 1h chaos run does not give cleanup an hour to complete.
const cleanupTimeout = 30 * time.Second

// run network emulation command, stop netem on timeout or abort
func runNetem(ctx context.Context, client netemClient, req *container.NetemRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// create new context with timeout for canceling

// wait for specified duration and then stop netem (where it applied) or stop on ctx.Done()
// use context.WithoutCancel so cleanup succeeds even if the parent ctx is canceled
// or if it inherited a deadline that has elapsed alongside stopCtx.
