package docker

import (
	ctr "github.com/alexei-led/pumba/pkg/container"
	ctypes "github.com/docker/docker/api/types/container"
	imagetypes "github.com/docker/docker/api/types/image"
)

// Containers list of containers
func Containers(containers ...ctypes.Summary) []ctypes.Summary {
	_ = "STUB: not implemented"

	// Response mock single container
	return nil
}

func Response(params map[string]any) ctypes.Summary {
	_ = "STUB: not implemented"
	return *new(ctypes.Summary)
}

// DetailsResponse mock container details response
func DetailsResponse(params map[string]any) ctypes.InspectResponse {
	_ = "STUB: not implemented"
	return *new(ctypes.InspectResponse)
}

// ImageDetailsResponse mock image response
func ImageDetailsResponse(params map[string]any) imagetypes.InspectResponse {
	_ = "STUB: not implemented"
	return *new(imagetypes.InspectResponse)
}

// NewTestContainer creates a Container directly from params for testing
func NewTestContainer(params map[string]any) *ctr.Container { _ = "STUB: not implemented"; return nil }

func lookupWithDefault(aMap map[string]any, key string, defaultValue any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// AsMap convert multiple arguments into map[string]any
func AsMap(args ...any) map[string]any { _ = "STUB: not implemented"; return nil }
