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
	CreateDt time.Time
 	// createDt time.Time
}

func AddShopUser(m ShopUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
// ShopUser

func GetAllShopUsers()([]ShopUser){
	o :=orm.NewOrm()
	var	shopUsers []ShopUser
	_,err :=	o.QueryTable(new (ShopUser)).All(&shopUsers)
	fmt.Println(shopUsers,err)
	return shopUsers
}

/**根据商城用户手机号获取用户手机**/
func (shopUser *ShopUser)GetShopUserByPhone(Phone string)(ShopUser,error){
	o :=orm.NewOrm()
	var user ShopUser=ShopUser{Phone:Phone}
	err :=o.Read(&user);if(err==nil){
		return user,nil 
	}else{
		return ShopUser{},err
	}
}
// ShopUser.
// ShopUser.