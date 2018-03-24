package controllers

import (
	"time"
	// "time"
	// "json"
	"fmt"
	"encoding/json"
	"errors"
	"github.com/24wings/bangwei-api/models"
	"strconv"
	"strings"
	"regexp"
	"github.com/24wings/alidayu"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


type ShopUserController struct {
	beego.Controller
}


type ApiCloudKeyScret struct{
	 ApiKey string
	 ApiScret string
}

var KeyScret =ApiCloudKeyScret{ApiKey:"LTAIcMnaxxUG7dbk",ApiScret: "VhNgQZrGYz7dXpiCUS8r36mbLgy6db"}



// URLMapping ...
func (c *ShopUserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	// c.Mapping("Signin",c.Sign)
}


// Post ...
// @Title Post
// @Description create Users
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 201 {int} models.Users
// @Failure 403 body is empty
// @router / [post]
func (c *ShopUserController) Post() {
	var v models.Users
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	// c.par
	 _, err := models.AddUsers(&v);if err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else { 
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Users by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Users
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ShopUserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetUsersById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Users
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Users
// @Failure 403
//  @router / [get]
func (c *ShopUserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUsers(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}


// Delete ...
// @Title Delete
// @Description delete the Users
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ShopUserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteUsers(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *ShopUserController) ShopUserSignin(){
	loginShopUser := models.ShopUser{}
	json.Unmarshal(c.Ctx.Input.RequestBody,&loginShopUser)
	c.ParseForm(&loginShopUser)
	// fmt.Println("Phone:",Phone)

	v,err :=	models.ShopUserService.GetShopUserByPhone(loginShopUser.Phone); if err == nil{
		if (v.Password== loginShopUser.Password){	
			c.Data["json"] =ShopUserResponse{Ok:true,Data:v}
		}else{
			c.Data["json"]=ErrorResponse{Ok:false,Data:"用户名或密码错误"}
	}
	}else{
		c.Data["json"]=ErrorResponse{Ok:false,Data:"用户名不存在"}	
	}
	   c.ServeJSON()
	   return
}


func (c *ShopUserController) ShopUserSignup(){
	AuthCode := c.GetString("AuthCode")
	var newShopUser  models.ShopUser
	c.ParseForm(&newShopUser);
	json.Unmarshal(c.Ctx.Input.RequestBody, &newShopUser)
	fmt.Println(newShopUser.Phone,AuthCode)
	authPassword,_ :=	regexp.MatchString("^1[0-9]{10}$",newShopUser.Phone)

	if(authPassword){
		nowStr := time.Now().Format("20060102")
		_,err :=	models.ShopUserService.GetShopUserByPhone(newShopUser.Phone); if err != nil  {
		success,detail,_	:=alidayu.QueryDetail(newShopUser.Phone, "邦为科技",nowStr,"SMS_127158851",KeyScret.ApiKey,KeyScret.ApiScret); if (success==true && detail.OutId==newShopUser.AuthCode){
			newUser,err	:=models.ShopUserService.AddShopUser(&newShopUser); if err==nil{
				c.Data["json"] =NumberResponse{Ok:true,Data:newUser}	
		}else{
			c.Data["json"]=ErrorResponse{Ok:false,Data:err.Error()}
		}
			
			}else{
				c.Data["json"]=ErrorResponse{Ok:false,Data:"验证码错误"}
		}
		fmt.Println(detail)
	}else{
			c.Data["json"]=ErrorResponse{Ok:false,Data:"该手机已经注册"}
		}
	}else{
		c.Data["json"]=ErrorResponse{Ok:false,Data:"请输入正确的手机号"}
	}
   	c.ServeJSON()
}

func (c *ShopUserController) ForgotShopUserPassword(){
	Phone :=c.GetString("Phone")
	// Password :=c.GetString("Password")
	NewPassword := c.GetString("NewPassword")
	AuthCode :=c.GetString("AuthCode")
	checkPhone,_  := regexp.MatchString("^1[0-9]{9,}$",Phone)
	// 检查用户格式
	if(checkPhone){
		nowStr :=time.Now().Format("20060102")
		
		user,err	 :=models.ShopUserService.GetShopUserByPhone(Phone); if(err ==nil){
		success,lastAuthCode,queryErr	:=alidayu.QueryDetail(Phone, "邦为科技",nowStr,"SMS_127158851",KeyScret.ApiKey,KeyScret.ApiScret); if(success && lastAuthCode.OutId==AuthCode){
			user.Password=NewPassword
			updateErr :=	models.ShopUserService.UpdateShopUserById(&user);if (updateErr==nil){
			c.Data["json"]=ErrorResponse{Ok:true,Data:"修改密码成功"}
		
		}else{
			c.Data["json"]=ErrorResponse{Ok:false,Data:err.Error()}
			fmt.Println("success...",success,lastAuthCode,queryErr)
		}
			
}else{
				c.Data["json"]=ErrorResponse{Ok:false,Data:"验证码错误"}
			}
		}else{
			c.Data["json"]=ErrorResponse{Ok:false,Data:"用户不存在"}
		}
	}else{
	c.Data["json"]=ErrorResponse{Ok:false,Data:"请输入正确的手机号"}
}
c.ServeJSON()
}


	
func (c *ShopUserController) SendMessage(){
	Phone := c.GetString("Phone")
	checkPhone,_	:=regexp.MatchString("^1[0-9]{10}$",Phone)
	if(checkPhone==false){
	c.Data["json"]=ErrorResponse{Ok:false,Data:"手机号码不合法"}
}else{
	success, resp,_ := alidayu.SendSMS(Phone, "邦为科技", "SMS_127158851", `1234`,KeyScret.ApiKey,KeyScret.ApiScret)
	c.Data["json"]=ErrorResponse{Ok:success,Data:resp}
}
	c.ServeJSON()
	return
}

func (c *ShopUserController) GetAllShopUsers(){
	o :=orm.NewOrm()	
	var users []models.ShopUser
	num,_	:=o.QueryTable(new (models.ShopUser)).All(&users)
		fmt.Println(num)
		c.Data["json"]=users
		c.ServeJSON()

}
