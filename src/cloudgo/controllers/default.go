package controllers

import (
	"github.com/astaxie/beego"
)

//初始页控制器
type MainController struct {
	beego.Controller
}

//应用控制器
type AppController struct {
	beego.Controller
}

//收到GET请求时执行的操作
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *AppController) Get() {
	c.Layout = "puzzle.html"
	c.TplName = "index.tpl"
}
