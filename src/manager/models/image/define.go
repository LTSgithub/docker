package image

import "github.com/docker/docker/api/types"

type PullRequest struct {
	Name string `json:"name"`
}
type PullResponse struct {
}

type DeleteRequest struct {
	Type    int    `json:"type"` // 0 通过id删除   1，通过tag删除
	TagName string `json:"tag_name"`
	ImageId string `json:"image_id"`
}
type DeleteResponse struct {
}

type ListRequest struct {
	FilterComplex string `json:"filter_complex"`
}
type ListResponse struct {
	TotalCount int      `json:"total_count"`
	List       []types.ImageSummary `json:"list"`
}
