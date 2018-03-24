package routers

import (
	"github.com/24wings/bangwei-api/controllers"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"

)

func init() {
	// 这段代码放在router.go文件的init()的开头
    beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// AllowAllOrigins:  true,
		AllowOrigins:[]string{"http://localhost:4200"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        AllowCredentials: true,
	}))
	
	beego.Router("/fenxiao/shop-user/signin",&controllers.ShopUserController{},"Post:ShopUserSignin")
	beego.Router("/fenxiao/shop-user/signup",&controllers.ShopUserController{},"Post:ShopUserSignup")
	beego.Router("/fenxaio/fenxaio-user/signup",&controllers.FenxiaoUserController{},"Post:FenxiaoUserSignup")
	beego.Router("/fenxiao/shop-user/forgot-password",&controllers.ShopUserController{},"Put:ForgotShopUserPassword")
	beego.Router("/fenxiao/user/user-auth-code",&controllers.ShopUserController{},"Get:SendMessage")
	beego.Router("/fenxiaoadmin/allfenxiaousers",&controllers.FenxiaoAdminController{},"Get:AllFenxiaoUsers")
	
	beego.AutoRouter(&controllers.FenxiaoUserController{})
	// beego.AutoRouter(&controllers.FenxiaoAdminController{})
	
}
