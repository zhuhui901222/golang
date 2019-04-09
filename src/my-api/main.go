package main

import (
	_ "my-api/routers"

	"github.com/astaxie/beego"
)

var MysqlPool int =434

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
