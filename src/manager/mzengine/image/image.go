package image

import (
	"context"
	"scode"
	"status"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	)

func List() ([]types.ImageSummary, error) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	imageList, err := cli.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	return imageList, nil
}

func Pull(tag string) (error) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	out, err := cli.ImagePull(ctx,tag, types.ImagePullOptions{})
	if err != nil {
		return status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}
	defer out.Close()

	return nil
}


func Delete(id string) error {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	_, err = cli.ImageRemove(ctx, id, types.ImageRemoveOptions{Force: true, PruneChildren: true})
	if err != nil {
		return status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	return nil
}
