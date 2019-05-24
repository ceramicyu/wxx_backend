package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	MainController

}
type User struct {
	Id int `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}
type GameList struct {
	GameId int `db:"game_id" json:"game_id"`
	GameName string `db:"game_name" json:"game_name"`
	GameIcon string `db:"game_icon" json:"game_icon"`
}


func (c *UserController) Post() {
	switch c.Input().Get("ac"){
		case "login":
		   c.Login()
	    case "gamelist":
			c.Gamelist()
		default:
			c.Login()
	}

}

func (c *UserController)Login(){
	username:=c.Input().Get("account")
	password:=c.Input().Get("password")
	o:=orm.NewOrm()
	r:=o.Raw("select * from game_user where username = ?",username)
	user:=&User{}
	err:=r.QueryRow(&user)
	if err!=nil{
		c.JCode(1,"用户未注册",user)
	}
	if password!=user.Password{
		c.JCode(1,"密码不正确",user)
	}
	c.JCode(0,"登录成功",user)
}


func (c *UserController)Gamelist(){
fmt.Println("aaaaaaaaaaa")
	o:=orm.NewOrm()
	r:=o.Raw("select * from game_list ")
	gl:=[]*GameList{}
	_,err:=r.QueryRows(&gl)
	if err!=nil{
		c.JCode(1,"未找到游戏列表",gl)
	}
	c.JCode(0,"登录成功",gl)
}