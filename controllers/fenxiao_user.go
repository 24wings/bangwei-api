package controllers
import (
	"fmt"
	// "encoding/json"
	"regexp"
	// "github.com/24wings/alidayu"
	"github.com/astaxie/beego"
	"github.com/24wings/bangwei-api/models"
)



type FenxiaoUserController struct {
	beego.Controller
}

func  (this *FenxiaoUserController)Test(){
		this.Data["json"] =ErrorResponse{Ok:true,Data:"成功"}
		this.ServeJSON()
}
// 分销用户登录
func (this *FenxiaoUserController) FenxiaoUserLogin(){
	var loginUser models.FenxiaoUser
	this.ParseForm(&loginUser)
 	checkPhone,_:=regexp.MatchString("^1[0-9]{10}$",loginUser.Phone);if checkPhone {
		 oldUser,err  :=	models.FenxiaoUserService.GetFenxiaoUserByPhone(loginUser.Phone);if err==nil && oldUser.Password==loginUser.Password {
			//  this.Data["json"]=NumberResponse{Ok:false,Data:oldUser.Id}
		 }else{
			 
			 this.Data["json"]=ErrorResponse{Ok:true,Data:"密码错误"}
		 }
	 }else{
		this.Data["json"]=ErrorResponse{Ok:false,Data:"请输入正确的手机号"}
	 }
	 this.ServeJSON()
	 return
}

// 分销用户注册
func (this *FenxiaoUserController) FenxiaoUserSignup(){
	AuthCode :=this.GetString("AuthCode")
	var	newFenxiaoUser	models.FenxiaoUser
	this.ParseForm(&newFenxiaoUser)
	// 检查手机号合法
	if (CheckPhone(newFenxiaoUser.Phone)){
		_,err :=models.FenxiaoUserService.GetFenxiaoUserByPhone(newFenxiaoUser.Phone);if err==nil{
			this.Data["json"]=ErrorResponse{Ok:false,Data:"该手机号已经注册"}
		}else{
		// 检查验证码正确
		success,detail :=QueryAuthCode(newFenxiaoUser.Phone); if (success&&detail.OutId==AuthCode){
			id,err :=	models.FenxiaoUserService.AddFenxiaoUser(&newFenxiaoUser) ;if err==nil{
				this.Data["json"]=NumberResponse{Ok:true,Data:id}		
		}else{
			this.Data["json"]=ErrorResponse{Ok:true,Data:"验证码错误"}
		}
		}else{
			this.Data["json"]=ErrorResponse{Ok:false,Data:err.Error()}
		}
		}
	}else{
		this.Data["json"]=ErrorResponse{Ok:false,Data:"请输入正确的手机号"}
	}
	this.ServeJSON();
	return;
}


// 分销用户忘记密码

func (this *FenxiaoUserController) FenxiaoUserForgotPassword(){
	NewPassword := this.GetString("Newpassword")
	AuthCode := this.GetString("AuthCode")
	Phone := this.GetString("Phone")

	oldUser ,err :=models.FenxiaoUserService.GetFenxiaoUserByPhone(Phone);if err==nil{
		success,detail := QueryAuthCode(Phone); if(success && detail.OutId==AuthCode ){
		fenxiaoUser :=	models.FenxiaoUser{Id:oldUser.Id,Password:NewPassword}
		_,err :=	models.FenxiaoUserService.UpdateFenxiaoUserById(&fenxiaoUser);if err==nil{
			this.Data["json"]=ErrorResponse{Ok:true,Data:"修改密码成功"}
		}else{
			this.Data["json"]=ErrorResponse{Ok:false,Data:err.Error()}
		}
		}else{
			this.Data["json"]=ErrorResponse{Ok:false,Data:"验证码错误"}
		}
	}else{
		this.Data["json"] = ErrorResponse{Ok:false,Data:"该手机号尚未注册"}
	}
	this.ServeJSON()
}

func (this *FenxiaoUserController) SubmitShopRecord(){
	picture :=this.GetString("Picture")
	var newShopRecord  models.ShopRecord
	// picture :=this.GetString("Picture")
	this.ParseForm(&newShopRecord)
	
	// json.Unmarshal(this.Ctx.Input.RequestBody,&newShopRecord)
	fmt.Println("picture",picture,newShopRecord.PictureId,newShopRecord.ReportUser)
	
	if(newShopRecord.ReportUser!=0&&newShopRecord.PictureId!=0){
	n,err:=	models.ShopRecordService.AddNewShopRecord(&newShopRecord);if err==nil{
		this.Data["json"]=NumberResponse{Ok:true,Data:n}
	}else{
		this.Data["json"]=ErrorResponse{Ok:true,Data:err.Error()}
	}}else{
		this.Data["json"]=ErrorResponse{Ok:false,Data:"参数不全"}
	}
	
	this.ServeJSON()
	
}