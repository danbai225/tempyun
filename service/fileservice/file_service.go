package fileservice

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
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
	var data struct {
		Files []entity.File `json:"files"`
	}
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
		if utils.Vdd(rms[i]) {
			return *utils.Err()
		}
		os.RemoveAll("files/" + path + rms[i])
	}
	return *r
}

//创建目录
func Mkdir(target string) entity.Rjson {
	if utils.Vdd(target) {
		return *utils.Err()
	}
	r := utils.Ok()
	os.MkdirAll("files/"+target, os.ModePerm)
	var data struct {
		File entity.File `json:"file"`
	}
	data.File = utils.FileInfo(target)
	r.Data = data
	return *r
}

//创建文件
func MkFile(target string) entity.Rjson {
	if utils.Vdd(target) {
		return *utils.Err()
	}
	r := utils.Ok()
	os.Create("files/" + target)
	var data struct {
		File entity.File `json:"file"`
	}
	data.File = utils.FileInfo(target)
	r.Data = data
	return *r
}

//重命名
func ReName(target string, name string) entity.Rjson {
	if utils.Vdd(target) || utils.Vdd(name) {
		return *utils.Err()
	}
	r := utils.Ok()
	os.Rename("files/"+target, "files/"+name)
	var data struct {
		File entity.File `json:"file"`
	}
	data.File = utils.FileInfo("files/" + name)
	r.Data = data
	return *r
}

//解压
func UnZip(dst, src string) (err error) {
	// 打开压缩文件，这个 zip 包有个方便的 ReadCloser 类型
	// 这个里面有个方便的 OpenReader 函数，可以比 tar 的时候省去一个打开文件的步骤
	zr, err := zip.OpenReader(src)
	defer zr.Close()
	if err != nil {
		return
	}

	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}

	// 遍历 zr ，将文件写入到磁盘
	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)

		// 如果是目录，就创建目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			// 因为是目录，跳过当前循环，因为后面都是文件的处理
			continue
		}

		// 获取到 Reader
		fr, err := file.Open()
		if err != nil {
			return err
		}

		// 创建要写出的文件对应的 Write
		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		n, err := io.Copy(fw, fr)
		if err != nil {
			return err
		}

		// 将解压的结果输出
		fmt.Printf("成功解压 %s ，共写入了 %d 个字符的数据\n", path, n)

		// 因为是在循环中，无法使用 defer ，直接放在最后
		// 不过这样也有问题，当出现 err 的时候就不会执行这个了，
		// 可以把它单独放在一个函数中，这里是个实验，就这样了
		fw.Close()
		fr.Close()
	}
	return nil
}
