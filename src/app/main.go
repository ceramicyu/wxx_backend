package main

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io"
	"net/http"
)

type Jcode struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func JCode(writer http.ResponseWriter,code int,msg string,data interface{}){
	res:=&Jcode{
		Code:code,
		Msg:msg,
		Data:data,
	}
    resp,_:=json.Marshal(res)
	 io.WriteString(writer,string(resp))
	return
}
func main(){

	beego.Run()
}