package routers

import (
	"cloudgo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//访问根页时，指向总控制器
	beego.Router("/", &controllers.MainController{})
	//访问次根页，指向应用控制器
	beego.Router("/puzzle", &controllers.AppController{})
}
