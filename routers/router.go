// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"jeedev-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/app",
			beego.NSInclude(
				&controllers.AppController{},
			),
		),
		beego.NSNamespace("/img",
			beego.NSInclude(
				&controllers.ImgController{},
			),
		),
		beego.NSNamespace("/area",
			beego.NSInclude(
				&controllers.AreaController{},
			),
		),
	)
	beego.AddNamespace(ns)

}
