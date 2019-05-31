package routers

import (
	"manager/controllers"

	"github.com/astaxie/beego"
)

func init() {

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/container",
			beego.NSInclude(
				&controllers.ContainerController{},
			),
		),
		beego.NSNamespace("/image",
			beego.NSInclude(
				&controllers.ImageController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
