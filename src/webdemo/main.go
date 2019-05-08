package main

import (
	"github.com/astaxie/beego/orm"
	_ "webdemo/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql","root:root@tcp(127.0.0.1:3306)/imooc")
	beego.Run()
}

