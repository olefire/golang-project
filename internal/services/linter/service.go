package linter

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"io"
)

type DockerClient struct {
	*client.Client
}

func NewClient(client *client.Client) *DockerClient {
	return &DockerClient{
		client,
	}
}

func (client *DockerClient) LintCode(ctx context.Context) io.ReadCloser {
	resp, err := client.ContainerCreate(ctx, &container.Config{
		Image: "local-dcycle-python-linter-image",
		Cmd:   []string{"--output-format=json", "./code"},
		Tty:   false,
	},
		&container.HostConfig{
			Binds: []string{
				fmt.Sprint("C:/Users/Zachar/Desktop/golang-project/code:/app/code"),
			},
		}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := client.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := client.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	options := types.ContainerRemoveOptions{
		Force:         true,
		RemoveVolumes: true,
	}

	err = client.ContainerRemove(ctx, resp.ID, options)

	if err != nil {
		panic(err)
	}

	return out
}
