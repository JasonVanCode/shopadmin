package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "shopadmin/initialize"
	_ "shopadmin/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
