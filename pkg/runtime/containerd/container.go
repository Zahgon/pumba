package containerd

import (
	"context"

	ctr "github.com/alexei-led/pumba/pkg/container"
	containerd "github.com/containerd/containerd/v2/client"
)

// toContainer converts a containerd container to the runtime-agnostic Container type.
// If all is false, only running containers (those with an active task) are included.
// Returns (container, skip, error) where skip=true means the container should be filtered out.
func toContainer(ctx context.Context, c containerd.Container, all bool) (*ctr.Container, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// resolveContainerName tries to extract a human-readable name from well-known
// container labels. Falls back to the container ID if no name label is found.
//
// Supported label sources (checked in priority order):
//   - Kubernetes: io.kubernetes.container.name (+ pod name + namespace)
//   - nerdctl:    nerdctl/name
//   - Docker:     com.docker.compose.service
func resolveContainerName(id string, labels map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}
