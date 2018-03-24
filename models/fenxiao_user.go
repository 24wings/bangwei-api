package models
import (
	"fmt"
	"time"
	// "fmt"
	"github.com/astaxie/beego/orm"
)

type FenxiaoUser struct {
	Id   int `orm:"auto"`
	Nickname string
	Phone string
	Password string
	CreateDt time.Time
	ReciveRegion string
	ReciveCity string
	ReciveArea string
	DetailAddress string
	AuthCode string
	
	WechatId string
	// ParentId int64  //用户上级
	Parent        *FenxiaoUser   `orm:"rel(one)"` 
	TotalMoney float64
	LessMoney float64
}

var FenxiaoUserService =FenxiaoUser{}

func (this *FenxiaoUser) GetFenxiaoUserByPhone(Phone string)(FenxiaoUser,error){
	o:=orm.NewOrm()
	var v FenxiaoUser;
	err	:= o.QueryTable(new(FenxiaoUser)).Filter("Phone",Phone).RelatedSel().One(&v);if err == nil {
		return v,nil
	}else{
	return  v,err

}
}

func (this *FenxiaoUser) AddFenxiaoUser(newFenxiaoUser *FenxiaoUser)( int64 , error){
	o :=orm.NewOrm()
	id,err :=o.Insert(newFenxiaoUser)
	return  id,err
} 
func (this *FenxiaoUser) UpdateFenxiaoUserById(m *FenxiaoUser)(int64 ,error){
	o := orm.NewOrm()
	v := FenxiaoUser{Id:m.Id}
	err := o.Read(&v);if err ==nil{
		 num,err :=	o.Update(m); 
		 return num,err
	}else{
		return 0,err
	}
}

func (this *FenxiaoUser) GetAllHadParent()([]FenxiaoUser,error){
	o := orm.NewOrm()
	// childFenxiaoUser :=FenxiaoUser{}
	var fenxiaoUsers []FenxiaoUser;
	num,err := o.QueryTable(new (FenxiaoUser)).RelatedSel("Parent").All(&fenxiaoUsers);
	fmt.Println(num,err)
	return fenxiaoUsers,err
	
}

func (this *FenxiaoUser) GetAllNoParent()([]FenxiaoUser,error){
	o :=orm.NewOrm()
	var fenxaioUsers  []FenxiaoUser
	_,err :=o.QueryTable(new(FenxiaoUser)).Filter("Parent__isnull",true).RelatedSel("Parent").All(&fenxaioUsers)
	// fmt.Println("noparent",err,fenxaioUser)
	
	
	return fenxaioUsers,err
}