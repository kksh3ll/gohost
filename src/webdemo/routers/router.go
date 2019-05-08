package routers

import (
    "github.com/astaxie/beego"
    "webdemo/controllers"
)

func init() {
    //beego.Router("/", &controllers.MainController{})

    beego.Include(&controllers.UserController{})
}
