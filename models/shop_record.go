package models

import (
	// "time"
	// "fmt"
	"github.com/astaxie/beego/orm"
)

var ShopRecordService = ShopRecord{}

type ShopRecord struct{
	Id int64
	PictureId int64
	Location string
	Telephone string
	AreaNum int64
	FloorNum int64
	WorkerNum int64
	MasterName  string
	MasterPhone string
	BussinessName string
	BussinessPhone string 
	ClockSystemBrand string 
	DaliySales string
	ReportUser int64
}

func (this *ShopRecord) AddNewShopRecord(m *ShopRecord)(int64 ,error){
	o := orm.NewOrm()
	n,err :=o.Insert(m)
	return n,err
}