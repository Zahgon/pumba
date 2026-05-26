package iptables

import (
	"net"

	"github.com/alexei-led/pumba/pkg/chaos"
	"github.com/alexei-led/pumba/pkg/chaos/cliflags"
	"github.com/alexei-led/pumba/pkg/container"
)

// RequestBase bundles the parsed iptables-level state shared by every
// per-action subcommand. Request carries the runtime fields (IPs, ports,
// duration, sidecar hint, dry-run); Iface and Protocol are kept separate so
// per-action parsers can assemble the iptables command prefix
// (`-I/-D INPUT -i <iface> [-p <proto>] …`); Limit is the --limit value
// consumed by the per-action ListNContainers call rather than by the runtime.
type RequestBase struct {
	Request  *container.IPTablesRequest
	Iface    string
	Protocol string
	Limit    int
}

// ParseRequestBase reads the iptables-level flags (--duration, --interface,
// --protocol, --source, --destination, --src-port, --dst-port,
// --iptables-image, --pull-image, --limit) from c and returns a RequestBase
// with the shared fields filled. Container, CmdPrefix and CmdSuffix on
// Request are left zero — each per-action Run sets them per iteration.
//
// c must be the iptables parent context. Per-action parsers pass c.Parent().
func ParseRequestBase(c cliflags.Flags, gp *chaos.GlobalParams) (*RequestBase, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateCIDRList(list []string) ([]*net.IPNet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
