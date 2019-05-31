package controllers

import (
	"manager/models/image"
)

type ImageController struct {
	Controller
}

// @router /pull		[post]
func (c *ImageController) Pull() {

	req := &image.PullRequest{}
	if err := c.parseRequestForm(req); err != nil {
		c.setOutput(nil, err)
		return
	}

	c.Controller.setOutput(image.Pull(req))
}

// @router /delete		[post]
func (c *ImageController) Delete() {

	req := &image.DeleteRequest{}
	if err := c.parseRequestForm(req); err != nil {
		c.setOutput(nil, err)
		return
	}

	c.Controller.setOutput(image.Delete(req))
}

// @router /list		[post]
func (c *ImageController) List() {

	req := &image.ListRequest{}
	if err := c.parseRequestForm(req); err != nil {
		c.setOutput(nil, err)
		return
	}

	c.Controller.setOutput(image.List(req))
}
