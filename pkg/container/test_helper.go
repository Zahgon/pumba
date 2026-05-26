package container

import (
	"io"
)

// CreateTestContainers create test container
func CreateTestContainers(count int) []*Container { _ = "STUB: not implemented"; return nil }

// CreateLabeledTestContainers generate test containers with labels
func CreateLabeledTestContainers(count int, labels map[string]string) []*Container {
	_ = "STUB: not implemented"
	return nil
}

// Wrap wraps a given text reader with a ReadCloser
func Wrap(text string) io.ReadCloser { _ = "STUB: not implemented"; return *new(io.ReadCloser) }

// DockerAPIResponse docker container api response body
type DockerAPIResponse struct {
	Container string
	File      string
	Status    string
}
