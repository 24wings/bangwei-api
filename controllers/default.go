package controllers

import (
	"github.com/astaxie/beego"
	"github.com/24wings/bangwei-api/models"
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