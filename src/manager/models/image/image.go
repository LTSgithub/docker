package image

import (
				"scode"
	"status"

	engine_image "manager/mzengine/image"

	"strings"

	"github.com/astaxie/beego"
		)

func Pull(req *PullRequest) (*PullResponse, error) {
	resp := &PullResponse{}

	if req.Name == "" {
		beego.Error("docker name is empty")
		return nil,status.NewStatusDesc(scode.ScodeManagerCommonParameterError,"docker name is empty")
	}

	if err := engine_image.Pull(req.Name);err != nil {
		beego.Error(err)
		return nil,status.NewStatusDesc(scode.ScodeManagerCommonParameterError,err.Error())
	}

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


	imageList, err := engine_image.List()
	if err != nil {
		beego.Error(err)
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	resp.List = imageList
	resp.TotalCount = len(resp.List)

	return resp, nil
}
