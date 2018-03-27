package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	connectToDb()
	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
	orm.RegisterModel(new(ShopUser))
	orm.RegisterModel(new(FenxiaoUser))
	orm.RegisterModel(new(ShopRecord))

}
