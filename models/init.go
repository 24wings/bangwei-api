package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	// time.Date()
	orm.RegisterModel(new(ShopUser))
	orm.RegisterModel(new(FenxiaoUser))
	
}