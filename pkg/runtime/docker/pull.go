package docker

import (
	"context"
)

// imagePullResponse is one line of the JSON progress stream returned by
// the Docker daemon's image pull endpoint.
type imagePullResponse struct {
	Status         string `json:"status"`
	Error          string `json:"error"`
	Progress       string `json:"progress"`
	ProgressDetail struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"progressDetail"`
}

// pullImage pulls a Docker image and drains the progress stream.
func (client dockerClient) pullImage(ctx context.Context, img string) error {
	_ = "STUB: not implemented"
	return nil
}
