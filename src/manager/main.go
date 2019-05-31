package main

import (
	_ "manager/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.AutoRender = false
	beego.Run()
}
