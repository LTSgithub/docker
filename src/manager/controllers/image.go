package controllers

import (
	"manager/models/image"
)

type ImageController struct {
	Controller
}

// @Title 下载镜像
// @router /pull		[post]
func (c *ImageController) Pull() {

	req := &image.PullRequest{}
	if err := c.parseRequestForm(req); err != nil {
		c.setOutput(nil, err)
		return
	}

	c.Controller.setOutput(image.Pull(req))
}

// @Title 下载镜像
// @router /delete		[post]
func (c *ImageController) Delete() {

	req := &image.DeleteRequest{}
	if err := c.parseRequestForm(req); err != nil {
		c.setOutput(nil, err)
		return
	}

	c.Controller.setOutput(image.Delete(req))
}

// @Title 获取镜像列表
// @router /list		[post]
func (c *ImageController) List() {

	req := &image.ListRequest{}
	if err := c.parseRequestForm(req); err != nil {
		c.setOutput(nil, err)
		return
	}

	c.Controller.setOutput(image.List(req))
}
