package main

import (
	"github.com/astaxie/beego"
	"tempyun/filter"
	_ "tempyun/routers"
	"tempyun/service/userservice"
)

func main() {
	filter.Ymfilter()
	filter.GetFilter()
	initPath()
	beego.Run()
}

func initPath()  {
	users:=userservice.AllUser()
	for i:=0;i<len(users);i++{
		beego.SetStaticPath("/"+users[i].Username,"files/"+users[i].Username)
	}
}
