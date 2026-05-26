package containerd

import (
	"fmt"
	"os"
)

// cgroupReader reads the cgroup file for a process. Overrideable in tests.
var cgroupReader = func(pid uint32) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("/proc/%d/cgroup", pid))
}

// isSystemdCgroup returns true if the cgroup parent path uses the systemd driver.
// Mirrors containerd CRI's heuristic: systemd cgroup paths end with ".slice".
func isSystemdCgroup(cgroupParent string) bool { _ = "STUB: not implemented"; return false }

// cgroupChildPath constructs a child cgroup path under the given parent.
// For systemd drivers (parent ends with ".slice"), it uses the systemd slice format:
//
//	"<slice>:pumba:<sidecarID>" → runc creates /<slice>/pumba-<sidecarID>.scope
//
// For cgroupfs drivers, it uses a simple path join: "<parent>/<sidecarID>".
func cgroupChildPath(cgroupParent, sidecarID string) string { _ = "STUB: not implemented"; return "" }

// resolveCgroupPath parses /proc/<pid>/cgroup and returns the target's cgroup path
// and its parent directory. On cgroups v2, expects "0::<path>". On cgroups v1,
// falls back to the "memory" subsystem path.
func resolveCgroupPath(pid uint32) (cgroupPath, cgroupParent string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// cgroups v2: hierarchy 0, empty subsystem list

// cgroups v1: look for memory subsystem as representative
