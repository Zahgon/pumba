package chaos

import (
	"context"

	"github.com/alexei-led/pumba/pkg/container"
)

// ContainerAction applies a chaos action to a single target container.
// In serial runs, the first error stops iteration. In parallel runs, every
// closure runs to completion regardless of errors; errgroup.Wait returns
// the first reported error.
type ContainerAction func(ctx context.Context, c *container.Container) error

// RunOnContainers lists running containers matching gp.{Names,Pattern,Labels}
// (capped by limit), optionally narrows to a single random pick when random
// is true, then invokes fn for each container. parallel selects between
// errgroup fanout (true) and a sequential for-loop (false). Returns nil when
// no containers match — same warning the per-action loops used to log.
//
// The helper takes container.Lister rather than the per-action narrow client
// interface so it stays domain-agnostic; every action's client embeds Lister.
//
// Example:
//
//	return chaos.RunOnContainers(ctx, n.client, n.gp, n.limit, random, true,
//	    func(ctx context.Context, c *container.Container) error {
//	        netemCtx, cancel := context.WithCancel(ctx)
//	        defer cancel()
//	        req := *n.req
//	        req.Container = c
//	        req.Command = netemCmd
//	        return runNetem(netemCtx, n.client, &req)
//	    })
func RunOnContainers(
	ctx context.Context,
	lister container.Lister,
	gp *GlobalParams,
	limit int,
	random, parallel bool,
	fn ContainerAction,
) error {
	_ = "STUB: not implemented"
	return nil
}

// RunOnContainersAll behaves like RunOnContainers but also includes stopped
// containers in the candidate set. Used by lifecycle.remove which can target
// non-running containers.
func RunOnContainersAll(
	ctx context.Context,
	lister container.Lister,
	gp *GlobalParams,
	limit int,
	random, parallel bool,
	fn ContainerAction,
) error {
	_ = "STUB: not implemented"
	return nil
}

func runOnContainers(
	ctx context.Context,
	lister container.Lister,
	gp *GlobalParams,
	limit int,
	all, random, parallel bool,
	fn ContainerAction,
) error {
	_ = "STUB: not implemented"
	return nil
}
