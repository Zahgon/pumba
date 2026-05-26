package iptables

import (
	"context"
	"time"

	"github.com/alexei-led/pumba/pkg/container"
)

const (
	ProtocolAny  = "any"
	ProtocolTCP  = "tcp"
	ProtocolUDP  = "udp"
	ProtocolICMP = "icmp"
)

// iptablesClient is the narrow interface needed by iptables commands.
type iptablesClient interface {
	container.Lister
	container.IPTables
}

// cleanupTimeout caps how long the iptables-cleanup sidecar cycle is allowed
// to run after abort or scheduled stop. Independent of --duration so a
// 1h chaos run does not give cleanup an hour to complete.
const cleanupTimeout = 30 * time.Second

// run iptables command, stop iptables on timeout or abort. The add/del prefix
// pair distinguishes the rule installation command (-I/-A/-N) from its mirror
// removal command (-D); both share the rest of the request fields.
func runIPTables(ctx context.Context, client iptablesClient, addReq, delReq *container.IPTablesRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// create new context with timeout for canceling

// wait for specified duration and then stop iptables (where it applied) or stop on ctx.Done()
// use context.WithoutCancel so cleanup succeeds even if the parent ctx is canceled
// or if it inherited a deadline that has elapsed alongside stopCtx.
