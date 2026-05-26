package docker

import (
	"context"

	ctr "github.com/alexei-led/pumba/pkg/container"
	ctypes "github.com/docker/docker/api/types/container"
	imagetypes "github.com/docker/docker/api/types/image"
)

// dockerInspectToContainer converts Docker inspect responses into a runtime-agnostic Container.
func dockerInspectToContainer(info ctypes.InspectResponse, img *imagetypes.InspectResponse) *ctr.Container {
	_ = "STUB: not implemented"
	return nil
}

// ListContainers returns a list of containers that match the given filter
func (client dockerClient) ListContainers(ctx context.Context, fn ctr.FilterFunc, opts ctr.ListOpts) ([]*ctr.Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (client dockerClient) listContainers(ctx context.Context, fn ctr.FilterFunc, opts ctypes.ListOptions) ([]*ctr.Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
