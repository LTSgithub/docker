package container

import (
	"github.com/docker/docker/client"
	"status"
	"scode"
	"github.com/docker/docker/api/types"
	"context"
)

func List() ([]types.Container, error){

	cli, err := client.NewClientWithOpts()
	if err != nil {
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	return containers,nil
}

func Inspect(id string)(*types.ContainerJSON, error)  {

	cli, err := client.NewClientWithOpts()
	if err != nil {
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	containerList, err := cli.ContainerInspect(context.Background(), id)
	if err != nil {
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	return containerList,nil
}