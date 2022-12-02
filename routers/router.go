package routers

import (
	"belog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/article/add", &controllers.AddArticleController{})
	beego.Router("/article/delete", &controllers.DetelArticlerController{})
	beego.Router("/article/update", &controllers.EditArticleController{})
}
