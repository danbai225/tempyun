package filter

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"strconv"
	"strings"
)

func Ymfilter() {
	var Filter = func(ctx *context.Context) {
		doamin := beego.AppConfig.String("doamin")
		if ctx.Input.Host() != doamin {
			req := httplib.Get(strings.Replace(ctx.Input.Site(), ctx.Input.Host(), doamin, -1) + ":" + strconv.Itoa(ctx.Input.Port()) + "/" + strings.Replace(ctx.Input.Host(), "."+doamin, "", -1) + ctx.Input.URL())
			bytes, _ := req.Bytes()
			ctx.ResponseWriter.Write(bytes)
		}
	}
	beego.InsertFilter("/*", beego.BeforeRouter, Filter)
}
func GetFilter() {
	var FilterUser = func(ctx *context.Context) {
		cookie := ctx.GetCookie("username")
		if true {
			doamin := beego.AppConfig.String("doamin")
			req := httplib.Get(strings.Replace("http://"+doamin+":"+strconv.Itoa(ctx.Input.Port())+"/"+cookie+ctx.Input.URL(), "/service/files", "", -1))
			bytes, _ := req.Bytes()
			ctx.ResponseWriter.Write(bytes)
		}
	}
	beego.InsertFilter("/service/files/*", beego.BeforeRouter, FilterUser)
}
