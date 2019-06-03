package container

import (
	engin_container "manager/mzengine/container"
	"scode"
	"status"

	"github.com/astaxie/beego"
)

func List(req *ListRequest) (*ListResponse, error) {
	resp := &ListResponse{}

	containerList, err := engin_container.List()
	if err != nil {
		beego.Error(err.Error())
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	resp.TotalCount = len(containerList)
	resp.List = containerList

	return resp, nil
}

func Inspect(req *InspectRequest) (*InspectResponse, error) {
	resp := &InspectResponse{}

	if req.Id == "" {
		beego.Error("id is empty")
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, "id is empty")
	}

	ret, err := engin_container.Inspect(req.Id)
	if err != nil {
		beego.Error(err)
		return nil, status.NewStatusDesc(scode.ScodeManagerCommonParameterError, err.Error())
	}

	resp.ContainerJSON = *ret

	return resp, nil
}
