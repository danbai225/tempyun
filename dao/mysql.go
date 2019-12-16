package dao

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	ip := beego.AppConfig.String("mysqlip")
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", user+":"+pass+"@tcp("+ip+":3306)/tempyun")
}

func Getcon() orm.Ormer {
	return orm.NewOrm()
}
