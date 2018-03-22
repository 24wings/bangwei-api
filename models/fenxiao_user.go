package models
import (
	"time"
	// "fmt"
	// "github.com/astaxie/beego/orm"
)

type FenxiaoUser struct {
	Id   int64 `orm:"auto"`
	// Name int
	Phone string
	Password string
	Nickname string
	CreateDt time.Time
 	// createDt time.Time
}