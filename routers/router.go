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
	beego.Router("/", &controllers.ShopUserController{},"Get:GetAllShopUsers")
	beego.Router("/users",&controllers.ShopUserController{},"Get:GetAll;POST:Post;Put:Put")
	beego.Router("/users/:id",&controllers.ShopUserController{},"*:GetOne")
	beego.Router("/fenxiao/shop-user/signin",&controllers.ShopUserController{},"Post:Signin")
	beego.Router("/fenxiao/shop-user/signup",&controllers.ShopUserController{},"Post:Signup")
	beego.Router("/fenxiao/user/forgot-password",&controllers.ShopUserController{},"Put:ForgotPassword")
	beego.Router("/fenxiao/user/user-auth-code",&controllers.ShopUserController{},"Get:SendMessage")

}
