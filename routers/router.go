package routers

import (
	"github.com/golangpkg/go-stat-reporter/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{}, "get:Index")
	beego.Router("/reporter/index", &controllers.MainController{}, "get:AdminIndex")
	//主要报表数据接口。
	beego.Router("/reporter/page/:pageId", &controllers.StatPageController{}, "get:PageHtml")
	beego.Router("/reporter/table/api", &controllers.StatPageController{}, "get:TableApi")
}
