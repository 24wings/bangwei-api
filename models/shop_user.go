package models
import (
	"time"
	"fmt"
	"github.com/astaxie/beego/orm"
)

 type ShopUser struct {
	Id   int64 `orm:"auto"`
	// Name int
	Phone string
	Password string
	Nickname string
	Sex int64
	// WechatId string
	Job string
	Industry string
	RecommandCode string
	createDt time.Time

	//
	AuthCode string
}


var ShopUserService=ShopUser{}

func (this ShopUser) AddShopUser(m *ShopUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
// ShopUser

func (this ShopUser)GetAllShopUsers()([]ShopUser){
	o :=orm.NewOrm()
	var	shopUsers []ShopUser
	_,err :=	o.QueryTable(new (ShopUser)).All(&shopUsers)
	fmt.Println(shopUsers,err)
	return shopUsers
}

/**根据商城用户手机号获取用户手机**/
func (this *ShopUser)GetShopUserByPhone(Phone string)(ShopUser,error){
	o :=orm.NewOrm()
	var user ShopUser=ShopUser{}
	err :=o.QueryTable(new(ShopUser)).Filter("Phone",Phone).RelatedSel().One(&user);if(err==nil){
		return user,nil 
	}else{
		return ShopUser{},err
	}
}


func (this ShopUser) UpdateShopUserById(m *ShopUser) (err error) {
	o := orm.NewOrm()
	v := ShopUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}