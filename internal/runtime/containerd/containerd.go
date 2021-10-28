package containerd

import (
	"context"
	"fmt"
	"os"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/oci"
	containerRuntime "github.com/danmrichards/container-runner/internal/runtime"
	"github.com/google/uuid"
	"github.com/opencontainers/runtime-spec/specs-go"
)

const rootFS = "rootfs.tar"

func init() {
	containerRuntime.Register(containerRuntime.ContainerD, New)
}

// ContainerD is a container runtime powered by the ContainerD SDK.
type ContainerD struct {
	client *containerd.Client
}

// Run implements containerRuntime.Runner.
func (c *ContainerD) Run(ctx context.Context, workloadDir, cmd string) (string, error) {
	id := uuid.NewString()

	f, err := os.Open(rootFS)
	if err != nil {
		return "", fmt.Errorf("open rootfs: %w", err)
	}

	// Import an image from the local guest filesystem.
	imgs, err := c.client.Import(ctx, f)
	if err != nil {
		return "", fmt.Errorf("import image: %w", err)
	}
	img := containerd.NewImage(c.client, imgs[0])

	// Create the container with the working directory mounted from the host.
	container, err := c.client.NewContainer(
		ctx,
		id,
		containerd.WithNewSnapshot(id, img),
		containerd.WithNewSpec(
			oci.WithImageConfig(img),
			oci.WithProcessCwd("/usr/local/bin"),
			oci.WithProcessArgs(cmd),
			oci.WithMounts([]specs.Mount{
				{
					Destination: "/usr/local/bin",
					Type:        "bind",
					Source:      workloadDir,
					Options:     []string{"bind", "rw"},
				},
			}),
		),
	)
	if err != nil {
		return "", fmt.Errorf("create container: %w", err)
	}

	task, err := container.NewTask(ctx, cio.NewCreator(cio.WithStdio))
	if err != nil {
		return "", fmt.Errorf("new task: %w", err)
	}
	defer task.Delete(ctx)

	if err = task.Start(ctx); err != nil {
		return "", fmt.Errorf("start task: %w", err)
	}

	return id, nil
}

// New returns an instantiated container runtime for Docker.
func New() (containerRuntime.Runner, error) {
	c, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		return nil, err
	}

	return &ContainerD{
		client: c,
	}, nil
}
