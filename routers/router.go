// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/doctorJE/users-go/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSRouter("/create", &controllers.UserController{}, "post:Create"),
			beego.NSRouter("/delete", &controllers.UserController{}, "post:Delete"),
			beego.NSRouter("/pwd/change", &controllers.UserController{}, "post:ChangePassword"),
			beego.NSRouter("/login", &controllers.UserController{}, "get:Login"),
		),
	)
	beego.AddNamespace(ns)
}
