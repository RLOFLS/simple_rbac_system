package main

import (
	_ "hello2/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//设置log
	logs.SetLogger(logs.AdapterFile, `{"filename":"./runtime/log/days.log","level":6,"daily":true,"maxdays":10,"color":true}`)
	//设置orm
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:12345678@/simple_oa_system?charset=utf8&loc=Asia%2FShanghai")
	beego.Run()
}

