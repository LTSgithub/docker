package container

import "github.com/docker/docker/api/types"

type InspectRequest struct {
	Uuid string `form:"uuid"`
}

type InspectResponse struct {
	Result string `json:"result"`
}

type ListRequest struct {
	FilterComplex string `json:"filter_complex"`
}
type ListResponse struct {
	TotalCount int               `json:"total_count"`
	List       []types.Container `json:"list"`
}
