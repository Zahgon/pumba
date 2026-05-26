package docker

import (
	"crypto/tls"

	ctr "github.com/alexei-led/pumba/pkg/container"
	dockerapi "github.com/docker/docker/client"
)

// NewClient returns a new Client instance which can be used to interact with the Docker API.
func NewClient(dockerHost string, tlsConfig *tls.Config) (ctr.Client, error) {
	_ = "STUB: not implemented"
	return *new(ctr.Client), nil
}

// NewAPIClient returns a bare Docker SDK client. Exposed so alternate runtimes
// (e.g. Podman via the Docker-compat socket) can reuse the HTTP/TLS setup.
func NewAPIClient(dockerHost string, tlsConfig *tls.Config) (*dockerapi.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewFromAPI wraps an existing Docker SDK client as a ctr.Client. Exposed so
// alternate runtimes can reuse the Docker implementation via embedding.
func NewFromAPI(api *dockerapi.Client) (ctr.Client, error) {
	_ = "STUB: not implemented"
	return *new(ctr.Client), nil
}

type dockerClient struct {
	containerAPI dockerapi.ContainerAPIClient
	imageAPI     dockerapi.ImageAPIClient
	systemAPI    dockerapi.SystemAPIClient
}

// Close is a no-op for the Docker client; the underlying HTTP connections are managed by the SDK.
func (client dockerClient) Close() error { _ = "STUB: not implemented"; return nil }
