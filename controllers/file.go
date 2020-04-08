package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"os"
	"path"
	"tempyun/entity"
	"tempyun/models"
	"tempyun/service/fileservice"
	"tempyun/utils"
)

type FileController struct {
	beego.Controller
}

// @router /file/zip [post]
func (c *MainController) Zip() {
	useri := c.GetSession("user")
	if useri == nil {
		c.Ctx.WriteString("权限不足!")
		return
	}
	user := useri.(models.User)
	username := user.Username
	file, information, err := c.GetFile("filepond")
	if err == nil {
		defer file.Close()
		filename := information.Filename
		if utils.Vdd(c.GetString("path")) {
			c.Ctx.WriteString("非法路径!")
			return
		}
		c.SaveToFile("filepond", "files/"+username+"/"+filename)
		fileservice.UnZip("files/"+username+c.GetString("path"), "files/"+username+"/"+filename)
		os.RemoveAll("files/" + username + "/" + filename)
	}
	c.Ctx.WriteString("ok!")
}

// @router /file/service [get,post]
func (c *MainController) Service() {
	rjson := entity.Rjson{}

	if utils.Vdd(c.GetString("target")) {
		c.Ctx.WriteString("非法路径!")
		return
	}
	useri := c.GetSession("user")
	if useri == nil {
		c.Ctx.WriteString("权限不足!")
		return
	}
	user := useri.(models.User)
	switch c.GetString("cmd") {
	case "init":
		rjson = fileservice.Init()
		break
	case "ls":
		rjson = fileservice.Ls(user.Username + c.GetString("target"))
		break
	case "rm":
		rjson = fileservice.Rm(user.Username, c.GetStrings("target[]"))
		break
	case "upload":
		file, information, err := c.GetFile("file")
		defer file.Close()
		filename := information.Filename
		if utils.Vdd(filename) {
			c.Ctx.WriteString("非法路径!")
			return
		}
		err = c.SaveToFile("file", path.Join("files/"+user.Username+c.GetString("target"), filename))
		if err != nil {
			rjson = *utils.Err()
		} else {
			rjson = *utils.Ok()
			var data struct {
				File entity.File `json:"file"`
			}
			data.File = utils.FileInfo(user.Username + c.GetString("target") + filename)
			rjson.Data = data
		}
		break
	case "mkdir":
		rjson = fileservice.Mkdir(user.Username + c.GetString("target"))
		break
	case "touch":
		rjson = fileservice.MkFile(user.Username + c.GetString("target"))
		break
	case "rename":
		rjson = fileservice.ReName(user.Username+c.GetString("target"), user.Username+c.GetString("name"))
		break
	}
	b, _ := json.Marshal(rjson)
	c.Ctx.WriteString(string(b))
}
