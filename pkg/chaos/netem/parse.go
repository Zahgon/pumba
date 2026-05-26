package netem

import (
	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
)

// ParseRequestBase reads the netem-level flags (--duration, --interface,
// --target, --egress-port, --ingress-port, --tc-image, --pull-image, --limit)
// from c and returns a *container.NetemRequest with the shared base fields
// filled, plus the --limit value (consumed by per-action ListNContainers calls
// rather than by the runtime). Container and Command are left zero — each
// per-action Run sets them per iteration.
//
// c must be the netem parent context. Per-action parsers pass c.Parent().
func ParseRequestBase(c cliflags.Flags, gp *chaos.GlobalParams) (*container.NetemRequest, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
