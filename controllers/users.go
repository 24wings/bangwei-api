package controllers

import (
	// "time"
	"fmt"
	"encoding/json"
	"errors"
	"github.com/24wings/bangwei-api/models"
	"strconv"
	"strings"
	"regexp"
	// "github.com/24wings/bangwei-api/libs/apicloud/smssdk"
	// "github.com/ltt1987/alidayu"

	"github.com/24wings/alidayu"
	"github.com/astaxie/beego"
)

//  UsersController operations for Users
type UsersController struct {
	beego.Controller
}

type ApiCloudKeyScret struct{
	 ApiKey string
	 ApiScret string
}

var KeyScret =ApiCloudKeyScret{ApiKey:"LTAIcMnaxxUG7dbk",ApiScret: "VhNgQZrGYz7dXpiCUS8r36mbLgy6db"}



// URLMapping ...
func (c *UsersController) URLMapping() {
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
func (c *UsersController) Post() {
	var v models.Users
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddUsers(&v); err == nil {
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
func (c *UsersController) GetOne() {
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
func (c *UsersController) GetAll() {
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

// Put ...
// @Title Put
// @Description update the Users
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 200 {object} models.Users
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UsersController) Put() {
	// idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.ParseInt(idStr, 0, 64)
	v  := models.Users{}

	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateUsersById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
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
func (c *UsersController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteUsers(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *UsersController) Signin(){
	fmt.Println("signin")
	Phone :=c.GetString("Phone");
	Password := c.GetString("Password")
	v,err :=	models.GetUserByPhone(Phone)
	
	if(err == nil){
		if (v!=nil && v.Password== Password){	
			c.Data["json"] =UserResponse{Ok:true,Data:v}
		}else{
			c.Data["json"]=ErrorResponse{Ok:true,Data:"用户名或密码错误"}
		}
	
	}else{
		c.Data["json"]=ErrorResponse{Ok:false,Data:err.Error()}
	}
	// fmt.Println("en")
	   c.ServeJSON()
	   return
}


func (c *UsersController) Signup(){
	Phone := c.GetString("Phone")
	Password :=c.GetString("Password")
	// AuthCode :=c.GetString("AuthCode")
	authPassword,_ :=	regexp.MatchString("^1[0-9]{10}$",Phone)
	if(authPassword){
	
	userOld,err :=	models.GetUserByPhone(Phone)
	
		if (userOld==nil || err !=nil ){
			newUser,err	:=models.AddUsers(&models.Users{Phone:Phone,Password:Password}); if err==nil{
				c.Data["json"] =newUser	
			}else{
				c.Data["json"]=ErrorResponse{Ok:false,Data:err.Error()}
		}
	}else{
			c.Data["json"]=ErrorResponse{Ok:false,Data:"该手机号已经注册+"}
		}
	}else{
		c.Data["json"]=ErrorResponse{Ok:false,Data:"请输入正确的手机号"}
	}
	   
	   
   	c.ServeJSON()
}

func (c *UsersController) ForgotPassword(){
	Phone :=c.GetString("Phone")
	// Password :=c.GetString("Password")
	NewPassword := c.GetString("NewPassword")
	AuthCode :=c.GetString("AuthCode")
	checkPhone,_  := regexp.MatchString("^1[0-9]{9,}$",Phone)
	// 检查用户格式
	if(checkPhone){
		user,err	 :=models.GetUserByPhone(Phone); if(err ==nil){
		// queryDt :=	time.Now().Format("20060102")
		// fmt.Println(queryDt);
		success,lastAuthCode,queryErr	:=alidayu.QueryDetail(Phone, "邦为科技","20180322","SMS_127158851",KeyScret.ApiKey,KeyScret.ApiScret); if(success){
			fmt.Println(lastAuthCode)
			c.Data["json"]=lastAuthCode
		}else{
			c.Data["json"]=ErrorResponse{Ok:false,Data:queryErr}
			fmt.Println("success...",success,lastAuthCode,queryErr)
		}
		if(success &&lastAuthCode.OutId==AuthCode){
			user.Password=NewPassword
			updateErr :=	models.UpdateUsersById(user);if (updateErr==nil){
			c.Data["json"]=ErrorResponse{Ok:true,Data:"修改密码成功"}
		
	}else{
			c.Data["json"]=ErrorResponse{Ok:false,Data:"旧密码错误"}
		}
}else{
				c.Data["json"]=ErrorResponse{Ok:false,Data:lastAuthCode.OutId}
			}
		}else{
			c.Data["json"]=ErrorResponse{Ok:false,Data:err.Error()}
		}
	}else{
	c.Data["json"]=ErrorResponse{Ok:false,Data:"请输入正确的手机号"}
}
c.ServeJSON()
}


	
func (c *UsersController) SendMessage(){
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
