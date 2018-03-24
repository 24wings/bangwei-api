package controllers

import (
	"time"
	"regexp"
	"github.com/astaxie/beego"
	"github.com/24wings/bangwei-api/models"
	"github.com/24wings/alidayu"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


type ErrorResponse struct{
	Ok bool
	Data string
}

type UserResponse struct{
	Ok bool
	Data *models.Users
}
type MsgResponse struct{
	Ok bool
}
type ShopUserResponse struct{
	Ok bool 
	Data models.ShopUser
}
type NumberResponse struct{
	Ok bool 
	Data int64
}


func CheckPhone(Phone string)bool{
	check,_ := regexp.MatchString("^1[0-9]{d}$",Phone)
	return check
}

func QueryAuthCode(Phone string)(bool,alidayu.SMSSendDetail){
	nowStr := time.Now().Format("20060102")
	success,detail,_ := 	alidayu.QueryDetail(Phone, "邦为科技",nowStr,"SMS_127158851",KeyScret.ApiKey,KeyScret.ApiScret)
	return success,detail
}
type FenxiaoUsersResponse struct{
	Ok bool
	Data []models.FenxiaoUser
}