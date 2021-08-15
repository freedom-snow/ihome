package main

import (
	beego "github.com/astaxie/beego"
	_ "snowy.space/ihome/routers"
)

func main() {
	beego.SetStaticPath("/download", "./conf")
	beego.Run()
}
