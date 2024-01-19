package service

import (
	"OJ_sandbox/constant"
	"OJ_sandbox/cqe"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

func CodeExecute(c *gin.Context) {
	// 判断请求是否具有风险

}

// RunCodeByDocker
// 操作docker执行代码
func RunCodeByDocker(c *gin.Context) {
	// 获取请求参数
	cmd := cqe.CodeRequestCmd{}
	if err := c.BindJSON(&cmd); err != nil {
		c.JSON(400, "请求参数失效！")
		return
	}

	// 判断请求参数
	if err := cmd.Validate(); err != nil {
		c.JSON(400, "无效参数")
		return
	}

	// 创建docker连接
	cli, err := ConnectDocker()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("docker 链接成功")
	}
	language := cmd.Language
	image := GetLanguageImage(language)
	config := &container.Config{
		Image: image,
		Tty:   true,
	}

	// 拉取镜像
	err = PullImage(cli, image)
	if err != nil {
		fmt.Println("拉取镜像失败：", err.Error())
	}

	// 创建容器
	containerId, err := CreateContainer(cli, config, nil, nil, "test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("new containerId", containerId)

	// 启动容器
	// todo

	// 删除容器
	err = DeleteContainer(cli, containerId)
	if err != nil {
		fmt.Println("delete fail: ", err.Error())
	}

	err = GetContainers(cli)
	if err != nil {
		fmt.Println(err)
	}

	// 删除镜像
	err = DeleteImage(cli, image)
	if err != nil {
		fmt.Println("delete fail :", err.Error())
	}

}

func GetLanguageImage(language string) string {
	// 根据语言返回相应的Docker镜像
	switch language {
	case "go":
		return constant.GoLanguageImage
	case "python":
		return constant.PythonLanguageImage
	case "java":
		return constant.JavaLanguageImage8
	default:
		return ""
	}
}

// ConnectDocker
// 链接docker
func ConnectDocker() (cli *client.Client, err error) {
	cli, err = client.NewClientWithOpts(client.WithAPIVersionNegotiation(), client.WithHost("tcp://localhost:2375"))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return cli, nil
}

// CreateContainer
// 创建容器
func CreateContainer(cli *client.Client, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (containerId string, err error) {
	ctx := context.Background()

	//创建容器
	resp, err := cli.ContainerCreate(ctx, config, hostConfig, networkingConfig, nil, containerName)
	if err != nil {
		fmt.Println(err.Error())
	}
	return resp.ID, nil
}

// PullImage
// 拉取 Docker 镜像
func PullImage(cli *client.Client, imageName string) error {
	_, err := cli.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
	if err != nil {
		return err
	}

	return nil
}

// DeleteImage
// 删除 Docker 镜像
func DeleteImage(cli *client.Client, imageID string) error {

	_, err := cli.ImageRemove(context.Background(), imageID, types.ImageRemoveOptions{})
	if err != nil {
		return err
	}

	fmt.Println("Image deleted:", imageID)
	return nil
}

// DeleteContainer
// 删除 Docker 容器
func DeleteContainer(cli *client.Client, containerID string) error {

	err := cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{Force: true})
	if err != nil {
		return err
	}

	fmt.Println("Container deleted:", containerID)
	return nil
}

// GetContainers
// 获取容器列表
func GetContainers(cli *client.Client) error {
	//All-true相当于docker ps -a
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		//fmt.Println(err)
		return err
	}
	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
	return nil
}