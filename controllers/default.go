package controllers

import (
	"github.com/TitanBenny/weixin"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type WXController struct {
	beego.Controller
}

func (c *WXController) Get() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter.ResponseWriter
	appConfig := beego.AppConfig.String
	weixin.Initialize(appConfig("originId"), appConfig("appId"), appConfig("appSecret"), appConfig("token"), appConfig("encodingAESKey"))
	weixin.HandleAccess(w, r)
}
