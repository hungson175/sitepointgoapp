package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

func (main *MainController) HelloSitePoint() {
	main.Data["Website"] = "SonPH Goooooopher"
	main.Data["Email"] = "sphamhung@gmail.com"
	main.Data["EmailName"] = "Pham Hung Son"
	main.Data["Id"] = main.Ctx.Input.Param(":id")
	main.TplNames = "default/hello-sitepoint.tpl"
}
