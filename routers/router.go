package routers

import (
	"github.com/golangpkg/go-stat-reporter/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/admin/index", &controllers.MainController{}, "get:AdminIndex")
	//
	beego.Router("/admin/page", &controllers.StatPageController{}, "get:PageHtml")
	beego.Router("/admin/table/api", &controllers.StatPageController{}, "get:TableApi")
}
