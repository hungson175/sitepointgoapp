package routers

import (
	"github.com/astaxie/beego"
	"github.com/hungson175/sitepointgoapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello-world/:id([0-9]+)", &controllers.MainController{}, "get:HelloSitePoint")
}
