package docker

import (
	"context"

	ctr "github.com/alexei-led/pumba/pkg/container"
	ctypes "github.com/docker/docker/api/types/container"
)

// StressContainer starts stress test on a container (CPU, memory, network, io)
func (client dockerClient) StressContainer(ctx context.Context, req *ctr.StressRequest) (*ctr.StressResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// stressContainerConfig builds the container and host config for a stress-ng container.
// cgroupPath is the target's cgroup base path resolved from ContainerInspect (may be empty).
// For inject-cgroup mode: when cgroupPath is known, uses --cgroup-path; otherwise falls back
// to --target-id + --cgroup-driver.
func stressContainerConfig(targetID string, stressors []string, img, driver, cgroupParent, cgroupPath string, injectCgroup bool) (ctypes.Config, ctypes.HostConfig) {
	_ = "STUB: not implemented"
	return *new(ctypes.Config), *new(ctypes.HostConfig)
}

// default child-cgroup mode: use --cgroup-parent with the resolved path

// stressContainerCommand executes a stress-ng command in a stress-ng Docker container
// in the target container's cgroup.
func (client dockerClient) stressContainerCommand(ctx context.Context, targetID string, stressors []string, img string, pull, injectCgroup bool) (string, <-chan string, <-chan error, error) {
	_ = "STUB: not implemented"
	return "", nil, nil, nil
}

// create stress-ng container

// attach to stress-ng container, capturing stdout and stderr

// AutoRemove fires only on container exit; a never-attached container
// is leaked unless we remove it explicitly.

// copy stderr and stdout from attached reader

// inspect stress-ng container

// get status of stress-ng command

// start stress-ng container running stress-ng in target container cgroup

// AutoRemove fires only on container exit; a never-started container
// is leaked unless we remove it explicitly.
