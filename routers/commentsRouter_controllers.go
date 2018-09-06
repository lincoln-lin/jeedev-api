package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["jeedev-api/controllers:AppController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AppController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AppController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AppController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AppController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AppController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AppController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AppController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AppController"],
		beego.ControllerComments{
			Method: "Img",
			Router: `/:sid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AreaController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AreaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AreaController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AreaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AreaController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AreaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AreaController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AreaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:AreaController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:AreaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:ImgController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:ImgController"],
		beego.ControllerComments{
			Method: "Code",
			Router: `/:sid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:ImgController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:ImgController"],
		beego.ControllerComments{
			Method: "Check",
			Router: `/check/:sid/:code`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:ImgController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:ImgController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/get/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["jeedev-api/controllers:ImgController"] = append(beego.GlobalControllerRouter["jeedev-api/controllers:ImgController"],
		beego.ControllerComments{
			Method: "Set",
			Router: `/set`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
