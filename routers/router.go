package routers

import (
	beego "github.com/astaxie/beego"
	"snowy.space/ihome/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
