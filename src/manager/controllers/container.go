package controllers

import (
	"manager/models/container"
)

type ContainerController struct {
	Controller
}

// @router /inspect		[post]
func (c *ContainerController) Inspect() {

	req := &container.InspectRequest{}
	if err := c.parseRequestForm(req); err != nil {
		c.setOutput(nil, err)
		return
	}

	c.Controller.setOutput(container.Inspect(req))
}

// @router /list		[post]
func (c *ContainerController) List() {

	req := &container.ListRequest{}
	if err := c.parseRequestForm(req); err != nil {
		c.setOutput(nil, err)
		return
	}

	c.Controller.setOutput(container.List(req))
}
