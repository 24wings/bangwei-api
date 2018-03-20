package main

import (
    _ "github.com/go-sql-driver/mysql"
	_ "github.com/24wings/bangwei-api/routers"
	"github.com/astaxie/beego"
)

func main() { 
	beego.Run()
}

  