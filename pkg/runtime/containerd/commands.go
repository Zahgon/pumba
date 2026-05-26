package containerd

import "net"

// buildNetemCommands constructs tc commands for applying network emulation.
// When IP/port filters are specified, creates a priority-based queueing hierarchy:
//
//	       1:   root qdisc (prio)
//	      / | \
//	    1:1 1:2 1:3    classes
//	     |   |   |
//	   10:  20:  30:   qdiscs
//	   sfq  sfq  netem
//	band 0   1    2
//
// Matching traffic is routed to band 2 (netem), all other traffic flows through sfq.
func buildNetemCommands(netInterface string, netemCmd []string, ips []*net.IPNet, sports, dports []string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// Simple case: apply netem directly on root qdisc
//nolint:mnd

// IP/port filter case: prio qdisc + sfq + netem + u32 filters
//nolint:mnd

// buildStopNetemCommands constructs tc commands to remove network emulation.
// When filters were used, removes the priority qdisc hierarchy; otherwise just deletes root netem.
func buildStopNetemCommands(netInterface string, hasFilters bool) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// buildIPTablesCommands constructs one iptables command per IP/port filter,
// matching Docker's behavior of issuing separate rules per filter element.
func buildIPTablesCommands(cmdPrefix, cmdSuffix []string, srcIPs, dstIPs []*net.IPNet, sports, dports []string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

//nolint:mnd

//nolint:mnd

//nolint:mnd

//nolint:mnd

// No filters: single command with just prefix + suffix
