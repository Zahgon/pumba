package util

import (
	"net"
	"regexp"
)

// reInterface accepts a network-interface identifier: must start with a
// letter and contain only letters, digits, '.', ':', '_', or '-'. Used as a
// shell-injection guard before names flow into `tc qdisc dev <name>` and
// `iptables -i <name>`.
var reInterface = regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9.:_-]*`)

// ValidateInterfaceName returns an error if name is not a safe network
// interface identifier per reInterface. Empty names are rejected.
func ValidateInterfaceName(name string) error { _ = "STUB: not implemented"; return nil }

// GetPorts will split the string of comma separated ports and return a list of ports
func GetPorts(ports string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Handle no port case

// verifyPort will make sure the port is numeric and within the correct range
func verifyPort(port string) error { _ = "STUB: not implemented"; return nil }

// ensure IP string is in CIDR notation
func cidrNotation(ip string) string { _ = "STUB: not implemented"; return "" }

// ParseCIDR Parse IP string to IPNet
func ParseCIDR(ip string) (*net.IPNet, error) { _ = "STUB: not implemented"; return nil, nil }
