package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["manager/controllers:ContainerController"] = append(beego.GlobalControllerRouter["manager/controllers:ContainerController"],
		beego.ControllerComments{
			Method: "Inspect",
			Router: `/inspect`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manager/controllers:ContainerController"] = append(beego.GlobalControllerRouter["manager/controllers:ContainerController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manager/controllers:ImageController"] = append(beego.GlobalControllerRouter["manager/controllers:ImageController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manager/controllers:ImageController"] = append(beego.GlobalControllerRouter["manager/controllers:ImageController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["manager/controllers:ImageController"] = append(beego.GlobalControllerRouter["manager/controllers:ImageController"],
		beego.ControllerComments{
			Method: "Pull",
			Router: `/pull`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
