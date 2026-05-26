package container

import (
	"context"
)

// ListOpts list options
type ListOpts struct {
	All    bool
	Labels []string
}

// list filter
type filter struct {
	Names   []string
	Pattern string
	Opts    ListOpts
}

func listContainers(ctx context.Context, client Lister, names []string, pattern string, labels []string, all bool) ([]*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RandomContainer select random container
func RandomContainer(containers []*Container) *Container { _ = "STUB: not implemented"; return nil }

//nolint:gosec

// ListNContainers list containers up to specified limit
func ListNContainers(ctx context.Context, client Lister, names []string, pattern string, labels []string, limit int) ([]*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListNContainersAll list containers up to specified limit, optionally including stopped containers
func ListNContainersAll(ctx context.Context, client Lister, names []string, pattern string, labels []string, limit int, all bool) ([]*Container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
