package controllers

import (
	"fmt"
	// "fmt"
	// "encoding/json"
	// "regexp"
	// "github.com/24wings/alidayu"
	"github.com/astaxie/beego"
	"github.com/24wings/bangwei-api/models"
)
type FenxiaoAdminController struct {
	beego.Controller
}



func (this *FenxiaoAdminController)AllFenxiaoUsers(){
	users,err	:=models.FenxiaoUserService.GetAllNoParent(); if err==nil{
		this.Data["json"]=FenxiaoUsersResponse{Ok:true,Data:users}
	}else{
		fmt.Println(users)
		this.Data["json"]=ErrorResponse{Ok:false,Data:"尚未有分销用户"}
	}
	this.ServeJSON()
}