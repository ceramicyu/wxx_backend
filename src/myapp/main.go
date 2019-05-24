package main

import (
	"github.com/astaxie/beego/orm"
	_ "myapp/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	orm.RegisterDataBase("default",
		"mysql",
		"root:qweqwe@tcp(127.0.0.1:3306)/game_server?charset=utf8", 30)
}
func main() {
	beego.Run()
}

