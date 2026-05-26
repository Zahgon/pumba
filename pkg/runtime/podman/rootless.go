package podman

import (
	"github.com/docker/docker/api/types/system"
)

// rootlessMarker is the substring Podman publishes in /info SecurityOptions
// when the API is serving a rootless connection.
const rootlessMarker = "name=rootless"

// detectRootless reports whether the Podman API is serving a rootless socket
// by scanning /info's SecurityOptions for the `name=rootless` marker.
func detectRootless(info *system.Info) bool { _ = "STUB: not implemented"; return false }

// rootlessError returns the user-facing error emitted when a chaos command
// that requires kernel privileges (netem, iptables, stress) is invoked
// against a rootless Podman socket. Message includes both macOS and Linux
// remediation hints.
func rootlessError(cmd, socketURI string) error { _ = "STUB: not implemented"; return nil }
