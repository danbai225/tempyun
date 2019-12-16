package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Username string `orm:"pk"`
	Password string
	Email    string
	Headurl  string
}
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}
