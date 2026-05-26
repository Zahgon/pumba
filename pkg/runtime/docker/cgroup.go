package docker

import (
	"context"
)

const (
	cgroupDriverSystemd  = "systemd"
	cgroupDriverCgroupfs = "cgroupfs"
)

// cgroupDriver queries the Docker daemon for its cgroup driver.
// Returns the driver name or empty string with error on failure.
func (client dockerClient) cgroupDriver(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// containerLeafCgroup returns the leaf cgroup directory name for a container
// based on the cgroup driver. On cgroupfs the leaf is the container ID; on
// systemd it is a scope unit named "docker-<id>.scope".
func containerLeafCgroup(targetID, driver string) string { _ = "STUB: not implemented"; return "" }

// inspectCgroupParent returns the target container's CgroupParent from inspect.
// Returns empty string when CgroupParent is not set (standalone Docker defaults)
// or when inspect fails.
func (client dockerClient) inspectCgroupParent(ctx context.Context, targetID string) string {
	_ = "STUB: not implemented"
	return ""
}

// defaultCgroupParent returns the default cgroup parent path based on the Docker
// daemon's cgroup driver when the target container has no explicit CgroupParent set.
func defaultCgroupParent(targetID, driver string) string { _ = "STUB: not implemented"; return "" }

// stressResolveDriver resolves the cgroup driver, parent, and target cgroup path
// for stress container setup. For default mode, cgroupParent is the resolved path
// for --cgroup-parent. For inject-cgroup mode, cgroupPath is the target's full
// cgroup path (if known) to pass as --cgroup-path to cg-inject.
func (client dockerClient) stressResolveDriver(ctx context.Context, targetID string, injectCgroup bool) (driver, cgroupParent, cgroupPath string, err error) {
	_ = "STUB: not implemented"
	// resolve the cgroup driver first — needed for correct leaf cgroup naming
	return "", "", "", nil
}

// try inspect anyway; if it yields a parent we can still proceed

// infer driver from parent path: systemd parents end with .slice

// For default mode, CgroupParent must be a value Docker accepts.
// systemd requires a valid slice name (*.slice); cgroupfs accepts any path.
