package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// MyDockerClient is struct of MyDockerClient
type MyDockerClient struct {
	cli *client.Client
}

// NewMyDockerClient is create MyDockerClient
func NewMyDockerClient() *MyDockerClient {
	t := new(MyDockerClient)

	// クライアントの作成
	// APIのバージョンを1.39までとする
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"))
	if err != nil {
		log.Fatal(err)
	}

	t.cli = cli
	return t
}

// Containers is show containert information
func (t *MyDockerClient) Containers() {
	containers, err := t.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}

// ServerVersion is show server version
func (t *MyDockerClient) ServerVersion() {
	v, err := t.cli.ServerVersion(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	msgYellow("サーババージョンの表示")
	fmt.Printf("Docker Version %s %s %s\n\n", v.Version, v.Os, v.Arch)
}

// LaunchAlpine is launch alpine container
func (t *MyDockerClient) LaunchAlpine() {

	ctx := context.Background()
	msgYellow("イメージの取得")
	reader, err := t.cli.ImagePull(
		ctx,
		"docker.io/library/alpine",
		types.ImagePullOptions{},
	)
	if err != nil {
		log.Fatal(err)
	}
	{
		scanner := bufio.NewScanner(reader)
		var last string
		for scanner.Scan() {
			var d struct {
				Status string
			}
			json.Unmarshal(scanner.Bytes(), &d)
			if last != d.Status {
				log.Printf("  Image Pull: %s\n", d.Status)
				last = d.Status
			}
		}
	}

	msgYellow("コンテナの作成")
	resp, err := t.cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: "alpine",
			Cmd:   []string{"echo", "Hello World"},
			Tty:   true,
		},
		nil, nil, "")
	if err != nil {
		log.Fatal(err)
	}

	msgYellow("コンテナの開始")
	err = t.cli.ContainerStart(
		ctx,
		resp.ID,
		types.ContainerStartOptions{},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("コンテナの開始完了を待機")
	statusCh, errCh := t.cli.ContainerWait(
		ctx,
		resp.ID,
		container.WaitConditionNotRunning,
	)
	select {
	case err := <-errCh:
		if err != nil {
			log.Fatal(err)
		}
	case <-statusCh:
	}

	msgYellow("ログの表示")
	out, err := t.cli.ContainerLogs(
		ctx,
		resp.ID,
		types.ContainerLogsOptions{ShowStdout: true},
	)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, out)

	msgYellow("コンテナの終了")
	err = t.cli.ContainerRemove(
		ctx,
		resp.ID,
		types.ContainerRemoveOptions{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
