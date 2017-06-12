package routers

import (
	"expanse/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wx", &controllers.WXController{})
}
