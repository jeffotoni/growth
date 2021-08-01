package tests

import (
	"context"
	"io"
	"os"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/require"
)

const serverPort = "8080"

type ShutdownFunc func() error

func setupServer(t *testing.T, workspace Workspace) ShutdownFunc {
	ctx := context.Background()
	name := slug.Make(t.Name())

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	require.NoError(t, err)

	tar, err := archive.TarWithOptions(workspace.directory, &archive.TarOptions{})
	require.NoError(t, err)
	defer tar.Close()

	buildRes, err := cli.ImageBuild(ctx, tar, types.ImageBuildOptions{Tags: []string{name}, Remove: true})
	require.NoError(t, err)
	defer buildRes.Body.Close()
	io.Copy(os.Stdout, buildRes.Body)

	// Remove previously created containers.
	cli.ContainerRemove(ctx, name, types.ContainerRemoveOptions{Force: true})

	// All the projects in this repository must be exposed through the port 8080.
	newport, err := nat.NewPort("tcp", serverPort)
	require.NoError(t, err)

	createRes, err := cli.ContainerCreate(ctx, &container.Config{
		Image: name,
		ExposedPorts: nat.PortSet{
			newport: {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			newport: []nat.PortBinding{{
				HostPort: serverPort,
			}},
		},
	}, nil, nil, name)
	require.NoError(t, err)

	err = cli.ContainerStart(ctx, createRes.ID, types.ContainerStartOptions{})
	require.NoError(t, err)

	out, err := cli.ContainerLogs(ctx, createRes.ID, types.ContainerLogsOptions{ShowStdout: true})
	require.NoError(t, err)
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	return func() error {
		return cli.ContainerRemove(ctx, createRes.ID, types.ContainerRemoveOptions{Force: true})
	}
}
