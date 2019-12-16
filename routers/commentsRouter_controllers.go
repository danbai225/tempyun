package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["tempyun/controllers:MainController"] = append(beego.GlobalControllerRouter["tempyun/controllers:MainController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tempyun/controllers:MainController"] = append(beego.GlobalControllerRouter["tempyun/controllers:MainController"],
        beego.ControllerComments{
            Method: "Service",
            Router: `/file/service`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tempyun/controllers:MainController"] = append(beego.GlobalControllerRouter["tempyun/controllers:MainController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tempyun/controllers:MainController"] = append(beego.GlobalControllerRouter["tempyun/controllers:MainController"],
        beego.ControllerComments{
            Method: "Pan",
            Router: `/pan`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tempyun/controllers:MainController"] = append(beego.GlobalControllerRouter["tempyun/controllers:MainController"],
        beego.ControllerComments{
            Method: "Test",
            Router: `/test`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
