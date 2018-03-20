package routers

import (
	"github.com/24wings/bangwei-api/controllers"
	
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/users",&controllers.UsersController{},"Get:GetAll;POST:Post;Put:Put")
	beego.Router("/users/:id",&controllers.UsersController{},"*:GetOne")
	beego.Router("/fenxiao/user-login",&controllers.UsersController{},"Post:Login")


}
