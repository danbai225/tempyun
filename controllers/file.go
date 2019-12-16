package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"path"
	"tempyun/entity"
	"tempyun/models"
	"tempyun/service/fileservice"
	"tempyun/utils"
)

type FileController struct {
	beego.Controller
}

// @router /file/service [get,post]
func (c *MainController) Service() {
	rjson := entity.Rjson{}
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
		file, information, err := c.GetFile("file")                                                   //返回文件，文件信息头，错误信息
		defer file.Close()                                                                            //关闭上传的文件，否则出现临时文件不清除的情况  mmp错了好多次啊
		filename := information.Filename                                                              //将文件信息头的信息赋值给filename变量
		err = c.SaveToFile("file", path.Join("files/"+user.Username+c.GetString("target"), filename)) //保存文件的路径。保存在static/upload中   （文件名）
		if err != nil {
			rjson = *utils.Err()
		} else {
			rjson = *utils.Ok()
			var data struct{ File entity.File `json:"file"` }
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
		rjson = fileservice.ReName(user.Username + c.GetString("target"),user.Username +c.GetString("name"))
		break
	}
	b, _ := json.Marshal(rjson)
	c.Ctx.WriteString(string(b))
}
