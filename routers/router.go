package routers

import (
	"Beego001/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/ ", &controllers.MainController{})
    beego.Router("/register",&controllers.FuncController{})
}
