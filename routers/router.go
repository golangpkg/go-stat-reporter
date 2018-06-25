package routers

import (
	"github.com/golangpkg/go-stat-reporter/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
