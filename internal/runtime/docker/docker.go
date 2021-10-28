package docker

import (
	"context"
	"fmt"
	"os"

	containerRuntime "github.com/danmrichards/container-runner/internal/runtime"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/google/uuid"
)

const rootFS = "rootfs.tar"

func init() {
	containerRuntime.Register(containerRuntime.Docker, New)
}

// Docker is a container runtime powered by the Docker SDK.
type Docker struct {
	client *client.Client
}

// Run implements containerRuntime.Runner.
func (d *Docker) Run(ctx context.Context, workloadDir, cmd string) (string, error) {
	imgRef := uuid.NewString()

	f, err := os.Open(rootFS)
	if err != nil {
		return "", fmt.Errorf("open rootfs: %w", err)
	}

	// Import an image from the local guest filesystem.
	if _, err = d.client.ImageImport(
		ctx,
		types.ImageImportSource{
			Source:     f,
			SourceName: "-",
		},
		imgRef,
		types.ImageImportOptions{},
	); err != nil {
		return "", fmt.Errorf("import image: %w", err)
	}

	// Create the container with the working directory mounted from the host.
	resp, err := d.client.ContainerCreate(
		ctx,
		&container.Config{
			Image: imgRef,
			Cmd:   strslice.StrSlice{cmd},
		},
		&container.HostConfig{
			NetworkMode: "host",
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: workloadDir,
					Target: "/usr/local/bin",
				},
			},
		},
		nil,
		nil,
		"",
	)
	if err != nil {
		return "", fmt.Errorf("create container: %w", err)
	}

	if err = d.client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", fmt.Errorf("start container: %w", err)
	}

	return resp.ID, nil
}

// New returns an instantiated container runtime for Docker.
func New() (containerRuntime.Runner, error) {
	c, err := client.NewClientWithOpts(
		client.FromEnv, client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, fmt.Errorf("docker client: %w", err)
	}

	return &Docker{
		client: c,
	}, nil
}
