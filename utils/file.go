package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"tempyun/entity"
	"time"
)

//获取文件修改时间 返回unix时间戳
func GetFileModTime(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}

//getFileSize get file size by path(B)
func DirSizeB(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

//getFileSize get file size by path(B)
func getFileSize(path string) int64 {
	if !exists(path) {
		return 0
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

//exists Whether the path exists
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func FileInfo(filename string) entity.File {
	stat, err := os.Stat("files/" + filename)
	file := entity.File{}
	if err == nil {
		file.Time = GetFileModTime("files/" + filename)
		file.Path = strings.Replace(filename, Upath(filename), "", -1)
		file.Name = stat.Name()
		if stat.IsDir() {
			file.Type = "dir"
			file.Mode = "40777"
		} else {
			file.Type = "file"
			file.Mode = "100666"
		}
		file.Read = true
		file.Write = true
		file.Isdir = stat.IsDir()
		file.Size, _ = DirSizeB("files/" + filename)
	}
	return file
}
func Upath(path string) string {
	index := strings.Index(path, "/")
	if index > 0 {
		return path[0:index]
	}
	return "/"
}
func CopyFile(source, dest string) bool {
	if source == "" || dest == "" {
		log.Println("source or dest is null")
		return false
	}
	//打开文件资源
	source_open, err := os.Open(source)
	//养成好习惯。操作文件时候记得添加 defer 关闭文件资源代码
	if err != nil {
		log.Println(err.Error())
		return false
	}
	defer source_open.Close()
	//只写模式打开文件 如果文件不存在进行创建 并赋予 644的权限。详情查看linux 权限解释
	dest_open, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	//养成好习惯。操作文件时候记得添加 defer 关闭文件资源代码
	defer dest_open.Close()
	//进行数据拷贝
	_, copy_err := io.Copy(dest_open, source_open)
	if copy_err != nil {
		log.Println(copy_err.Error())
		return false
	} else {
		return true
	}
}