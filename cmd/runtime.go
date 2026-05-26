package main

import (
	"crypto/tls"

	ctr "github.com/alexei-led/pumba/pkg/container"
	"github.com/alexei-led/pumba/pkg/runtime/containerd"
	"github.com/alexei-led/pumba/pkg/runtime/docker"
	"github.com/alexei-led/pumba/pkg/runtime/podman"
	"github.com/urfave/cli"
)

// Runtime client factories. Package-level vars so tests can swap them without
// requiring a real Docker/containerd/podman socket.
var (
	newDockerClient     = docker.NewClient
	newContainerdClient = containerd.NewClient
	newPodmanClient     = podman.NewClient
)

// createRuntimeClient constructs the container.Client for the runtime selected
// via --runtime. Extracted from before() to keep gocyclo under the 15 limit
// and to give unit tests a single function to exercise.
func createRuntimeClient(c *cli.Context) (ctr.Client, error) {
	_ = "STUB: not implemented"
	return *new(ctr.Client), nil
}

// tlsConfig still reads *cli.Context directly: it mixes flag reads with
// os.ReadFile/x509 helpers, so threading the adapter would add noise
// without payoff for the v3 migration this abstraction targets.

// tlsConfig translates the command-line options into a tls.Config struct
func tlsConfig(c *cli.Context) (*tls.Config, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:gosec

// Load CA cert

// Load client certificate
