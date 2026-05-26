package docker

import (
	"context"
	"time"

	ctr "github.com/alexei-led/pumba/pkg/container"
)

// sidecarRemoveTimeout bounds how long pumba will wait for ContainerRemove
// to reap an ephemeral tc/iptables sidecar after the caller's ctx cancels
// (e.g. SIGTERM). Podman's force-remove can take a few seconds on slow VMs.
const (
	sidecarRemoveTimeout  = 15 * time.Second
	sidecarInspectTimeout = 2 * time.Second
)

// removeSidecar force-removes an ephemeral tc/iptables sidecar container.
// Uses context.WithoutCancel with a short timeout so cleanup still runs
// when the caller's ctx was canceled by SIGTERM — otherwise pumba would
// leak the sidecar AND the rules it installed in the target's netns,
// because the caller early-returns on this error.
func (client dockerClient) removeSidecar(ctx context.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

func (client dockerClient) sidecarRemovalComplete(ctx context.Context, id string, removeErr error) bool {
	_ = "STUB: not implemented"
	return false
}

// runSidecar launches an ephemeral sidecar container that joins target's
// network namespace, runs argsList through `tool` (tc or iptables), and is
// force-removed on completion. Used by both netem and iptables paths.
func (client dockerClient) runSidecar(ctx context.Context, target *ctr.Container, argsList [][]string, img, tool string, pull bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Explicit Entrypoint/Cmd keeps the sidecar alive regardless of the
// image's default (e.g. nicolaka/netshoot defaults to zsh which exits
// immediately in detached mode). StopSignal: SIGKILL skips the
// SIGTERM-then-wait grace period on `rm -f`: tail as PID 1 ignores
// SIGTERM, which otherwise makes Podman wait the full 10 s StopTimeout
// before escalating (~tens of seconds per chaos cycle).

func (client dockerClient) pullSidecarImage(ctx context.Context, img, tool string) error {
	_ = "STUB: not implemented"
	return nil
}

// runSidecarExec creates and runs an exec inside the sidecar container,
// invoking `tool` (tc or iptables) with args. The exit code is inspected so
// that a non-zero status (e.g. tc rejecting bad args, iptables rule rejected
// by kernel) surfaces as an error instead of silent success.
func (client dockerClient) runSidecarExec(ctx context.Context, sidecarID, tool string, args []string) error {
	_ = "STUB: not implemented"
	return nil
}
