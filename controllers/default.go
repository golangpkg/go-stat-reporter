package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.Redirect("/admin/index", 302)
}

func (c *MainController) AdminIndex() {
	c.TplName = "index.html"
}
