package image

import (
	"context"
	"io"
	"os"

	"scode"
	"status"

	engine_image "manager/engine/image"

	"strings"

	"github.com/astaxie/beego"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Pull(req *PullRequest) (*PullResponse, error) {
	resp := &PullResponse{}

	//if req.Name == "" {
	//	beego.Error("name is empty")
	//	return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, "name is empty")
	//}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		beego.Error(err)
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	out, err := cli.ImagePull(ctx, req.Name, types.ImagePullOptions{})
	if err != nil {
		beego.Error(err)
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}
	defer out.Close()

	io.Copy(os.Stdout, out)

	return resp, nil
}

func Delete(req *DeleteRequest) (*DeleteResponse, error) {
	resp := &DeleteResponse{}

	if req.Type != 0 && req.Type != 1 {
		beego.Error("type candidate : 0 1")
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, "type candidate : 0 1")
	}

	imageList, err := engine_image.List()
	if err != nil {
		beego.Error(err)
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	// 根据tag name找到image的id
	idList := []string{}
	if req.Type == 0 {
		idList = append(idList, req.ImageId)
	} else {
		for i := 0; i < len(imageList); i++ {
			for j := 0; j < len(imageList[i].RepoTags); j++ {
				if strings.Contains(imageList[i].RepoTags[j], req.TagName) {
					idList = append(idList, imageList[i].ID)
				}
			}
		}

	}

	// 删除
	for i := 0; i < len(idList); i++ {
		if err := engine_image.Delete(idList[i]); err != nil {
			beego.Error(err)
		}
	}

	return resp, nil
}

func List(req *ListRequest) (*ListResponse, error) {
	resp := &ListResponse{}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts()
	if err != nil {
		beego.Error(err)
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	imageList, err := cli.ImageList(ctx, types.ImageListOptions{All: true})
	if err != nil {
		beego.Error(err)
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	for i := 0; i < len(imageList); i++ {
		beego.Debug("------------------tag  :%v", imageList[i].RepoTags)
	}

	return resp, nil
}
