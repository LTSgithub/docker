package container

import "github.com/docker/docker/api/types"

type InspectRequest struct {
	Id string `json:"id"`
}

type InspectResponse struct {
	types.ContainerJSON
}

type ListRequest struct {
	FilterComplex string `json:"filter_complex"`
}
type ListResponse struct {
	TotalCount int               `json:"total_count"`
	List       []types.Container `json:"list"`
}
