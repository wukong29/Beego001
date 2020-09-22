package controllers

import (
"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "2289476116@qq.com"
	c.TplName = "index.tpl"
}
