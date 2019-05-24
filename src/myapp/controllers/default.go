package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func(m MainController)JCode(code int,msg string,data interface{}){
	m.Data["json"] = map[string]interface{}{"code": code,"msg":msg,"data":data}
	m.ServeJSON()
	m.StopRun()
}
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
