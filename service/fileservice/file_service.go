package fileservice

import (
	"io/ioutil"
	"os"
	"tempyun/entity"
	"tempyun/utils"
)

//初始化 根目录
func Init() entity.Rjson {
	r := utils.Ok()
	var data struct {
		Config []string    `json:"config"`
		Root   entity.File `json:"root"`
	}
	data.Config = []string{}
	data.Root = utils.FileInfo("danbai")
	r.Data = data
	return *r
}

//输出目标目录
func Ls(target string) entity.Rjson {
	r := utils.Ok()
	var data struct{ Files []entity.File `json:"files"` }
	dir, err := ioutil.ReadDir("files/" + target)
	if err == nil {
		len := len(dir)
		for i := 0; i < len; i++ {
			files := append(data.Files, utils.FileInfo(target+dir[i].Name()))
			data.Files = files
		}
	}
	r.Data = data
	return *r
}

//移除文件或目录
func Rm(path string, rms []string) entity.Rjson {
	r := utils.Ok()
	for i := 0; i < len(rms); i++ {
		if utils.Vdd(rms[i]){
			return *utils.Err()
		}
		os.RemoveAll("files/" + path + rms[i])
	}
	return *r
}

//创建目录
func Mkdir(target string) entity.Rjson {
	if utils.Vdd(target){
		return *utils.Err()
	}
	r := utils.Ok()
	os.MkdirAll("files/"+target, os.ModePerm)
	var data struct{ File entity.File `json:"file"` }
	data.File = utils.FileInfo(target)
	r.Data = data
	return *r
}

//创建文件
func MkFile(target string) entity.Rjson {
	if utils.Vdd(target){
		return *utils.Err()
	}
	r := utils.Ok()
	os.Create("files/" + target)
	var data struct{ File entity.File `json:"file"` }
	data.File = utils.FileInfo(target)
	r.Data = data
	return *r
}

//重命名
func ReName(target string, name string) entity.Rjson {
	if utils.Vdd(target)||utils.Vdd(name){
		return *utils.Err()
	}
	r := utils.Ok()
	os.Rename("files/"+target, "files/"+name)
	var data struct{ File entity.File `json:"file"` }
	data.File = utils.FileInfo("files/" + name)
	r.Data = data
	return *r
}
