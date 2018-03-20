package routers

import (
	"github.com/24wings/bangwei-api/controllers"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"

)

func init() {
	// 这段代码放在router.go文件的init()的开头
    beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowAllOrigins:  true,
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        AllowCredentials: true,
    }))
	beego.Router("/", &controllers.MainController{})
	beego.Router("/users",&controllers.UsersController{},"Get:GetAll;POST:Post;Put:Put")
	beego.Router("/users/:id",&controllers.UsersController{},"*:GetOne")
	beego.Router("/fenxiao/user/signin",&controllers.UsersController{},"Post:Signin")
	beego.Router("/fenxiao/user/signup",&controllers.UsersController{},"Post:Signup")
	beego.Router("/fenxiao/user/forgot-password",&controllers.UsersController{},"Put:ForgotPassword")

}
