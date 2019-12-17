package userservice

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"reflect"
	"tempyun/dao"
	"tempyun/models"
	"tempyun/utils"
)

func AddUser(user models.User) bool {
	var o = dao.Getcon()
	user.Password = utils.Md5(user.Username + user.Password)
	_, err := o.Insert(&user)
	if err != nil {
		return false
	}
	os.Mkdir("./files/"+user.Username, os.ModePerm)
	utils.CopyFile("files/index.html", "files/"+user.Username+"/index.html")
	beego.SetStaticPath("/"+user.Username, "files/"+user.Username)
	return true
}
func UpdateUser(user models.User) bool {
	var o = dao.Getcon()
	user1 := models.User{Username: user.Username}
	if o.Read(&user1) == nil {
		varType := reflect.ValueOf(user)
		varType1 := reflect.ValueOf(&user1).Elem()
		for i := 1; i < 4; i++ {
			v := varType.Field(i).String()
			if v != "" {
				varType1.Field(i).SetString(v)
			}
		}
		if num, err := o.Update(&user1); err == nil {
			fmt.Println(num)
			if num > 0 {
				return true
			}
		}
	}
	return false
}
func GetUser(username string) models.User {
	var o = dao.Getcon()
	user := models.User{Username: username}
	err := o.Read(&user)
	if err != nil {
		user.Username = ""
	}
	return user
}

func AllUser() []models.User {
	var o = dao.Getcon()
	var users []models.User
	o.Raw("SELECT * FROM user").QueryRows(&users)
	return users
}
func VerifyUser(user models.User) (bool, models.User) {
	getUser := GetUser(user.Username)
	if getUser.Password == utils.Md5(user.Username+user.Password) {
		return true, getUser
	}
	return false, user
}
