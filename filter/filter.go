package filter

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"strconv"
	"strings"
)

//子域名转换
func Ymfilter() {
	var Filter = func(ctx *context.Context) {
		//跨域
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.Output.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
		ctx.Output.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
		ctx.Output.Header("Access-Control-Allow-Credentials", "true")
		doamin := beego.AppConfig.String("doamin")
		if ctx.Input.Host() != doamin {
			req := httplib.Get(strings.Replace(ctx.Input.Site(), ctx.Input.Host(), doamin, -1) + ":" + strconv.Itoa(ctx.Input.Port()) + "/" + strings.Replace(ctx.Input.Host(), "."+doamin, "", -1) + ctx.Input.URL())
			bytes, _ := req.Bytes()
			ctx.ResponseWriter.Header().Set("Content-Type",req.GetRequest().Header.Get("Content-Type"))
			ctx.ResponseWriter.Write(bytes)
		}
	}
	beego.InsertFilter("/*", beego.BeforeRouter, Filter)
}

//用户文件管理 路径转换过滤器
func GetFilter() {
	var FilterUser = func(ctx *context.Context) {
		user := ctx.GetCookie("username")
		if user != "" {
			doamin := beego.AppConfig.String("doamin")
			req := httplib.Get(strings.Replace("http://"+doamin+":"+strconv.Itoa(ctx.Input.Port())+"/"+user+ctx.Input.URL(), "/service/files", "", -1))
			bytes, _ := req.Bytes()
			ctx.ResponseWriter.Write(bytes)
		}
	}
	beego.InsertFilter("/service/files/*", beego.BeforeRouter, FilterUser)
}
