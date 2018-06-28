package main

import (
	_ "github.com/golangpkg/go-stat-reporter/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/golangpkg/go-stat-reporter/models"
	"os"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/context"
	"html/template"
)

func initDb() {
	//数据库注册。
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	db := beego.AppConfig.String("db")
	// 参数4(可选)  设置最大空闲连接
	maxIdle := 30
	// 参数5(可选)  设置最大数据库连接
	maxConn := 30
	conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", conn, maxIdle, maxConn)
	orm.Debug = true
	//同步 ORM 对象和数据库
	//这时, 在你重启应用的时候, beego 便会自动帮你创建数据库表。
	//orm.RunSyncdb("default", false, true)
}

func intiFilter() {
	//增加拦截器。
	var FilterUserByAdmin = func(ctx *context.Context) {
		url := ctx.Input.URL()
		pageName := ctx.Input.Query("pageId")
		//登录页面不过滤。
		if url == "/admin/userInfo/login" || url == "/admin/userInfo/logout" {
			return
		}
		logs.Info("########### url: ", url, ", pageName:", pageName, ",ParamsLen", ctx.Input.ParamsLen())
		//ctx.Input.SetData(urlTag, true)
		ctx.Input.SetData("PageList", &models.ConstantXmlPages.Pages)
	}
	beego.InsertFilter("/admin/*", beego.BeforeExec, FilterUserByAdmin)

}

//总体初始化。
func init() {
	//初始化数据库。
	initDb()
	//初始化filter
	intiFilter()
	//开启session。配置文件 配置下sessionon = true即可。
	beego.BConfig.WebConfig.Session.SessionOn = true
	//初始化xml数据模板。
	pwd, _ := os.Getwd()
	println("get pwd:", pwd)
	models.ReadXMLConfig(pwd + "/conf/pages.xml")
	println("######################")
}

//http://blog.xiayf.cn/2013/11/01/unescape-html-in-golang-html_template/
// 定义函数unescaped
func rawJs(x string) interface{}   { return template.JS(x) }
func rawHtml(x string) interface{} { return template.HTML(x) }

func main() {
	//增加自定义函数。
	beego.AddFuncMap("rawJs", rawJs)
	beego.AddFuncMap("rawHtml", rawHtml)
	//放到最后。
	beego.Run()
}
