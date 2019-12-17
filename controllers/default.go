package controllers

import (
	"github.com/astaxie/beego"
	"regexp"
	"tempyun/models"
	"tempyun/service/userservice"
)

type MainController struct {
	beego.Controller
}

// @router / [get]
func (c *MainController) Index() {
	c.TplName = "index.gohtml"
}

// @router /login [get,post]
func (c *MainController) Login() {
	if c.GetSession("user") != nil {
		c.Redirect("/tpan", 302)
		return
	}
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "login.gohtml"
		return
	}
	ok, user := userservice.VerifyUser(models.User{Username: c.GetString("username"), Password: c.GetString("password")})
	if ok {
		c.SetSession("user", user)
		c.Ctx.SetCookie("username", user.Username, 60*60)
		c.Ctx.SetCookie("beegosessionID", c.CruSession.SessionID(), 60*60)
		c.Redirect("/tpan", 302)
	} else {
		c.Data["Msg"] = "账号密码错误"
		c.TplName = "login.gohtml"
	}
}

// @router /reg [get,post]
func (c *MainController) Reg() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "reg.gohtml"
		return
	}
	pass:=c.GetString("password")
	username := c.GetString("username")
	if len(username)<1||len(pass)<4{
		c.Data["Msg"] = "长度太短哦"
		c.TplName = "reg.gohtml"
		return
	}
	if pass != c.GetString("vpassword") {
		c.Data["Msg"] = "密码不一致"
	}

	if ok, _ := regexp.MatchString("^[a-z0-9]{4,16}$", username); !ok {
		c.Data["Msg"] = "用户名为a-z 0-9 的4-16长度字符串"
	}
	user := userservice.GetUser(username)
	if user.Username == "" {
		email := c.GetString("email")
		us := models.User{Username: username, Password: pass, Email: email, Headurl: "http://www.gravatar.com/avatar/" + username + "?s=128&d=monsterid"}
		if userservice.AddUser(us) {
			c.Data["Msg"] = "注册成功"
		}
	} else {
		c.Data["Msg"] = "用户名已存在"
	}
	c.TplName = "reg.gohtml"
	return
}

// @router /tpan [get]
func (c *MainController) Pan() {
	if c.GetSession("user") == nil {
		c.Redirect("/login", 302)
	}
	c.Data["User"] = c.GetSession("user")
	c.TplName = "pan.gohtml"
}
