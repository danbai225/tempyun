package routers

import (
	"github.com/astaxie/beego"
	"tempyun/controllers"
)

func init() {
	beego.Include(&controllers.MainController{}, &controllers.FileController{})
}
