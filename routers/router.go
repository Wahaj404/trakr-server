// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"trakr-server/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			// get all users
			beego.NSInclude(&controllers.UserController{}),

			beego.NSRouter("/",
				&controllers.UserController{},
				"get:GetAll",
			),

			// add new user
			beego.NSRouter("/", &controllers.UserController{}, "post:Post"),

			// update an existing user
			beego.NSRouter("/:uid", &controllers.UserController{}, "put:Put"),

			// delete a user
			beego.NSRouter("/:uid", &controllers.UserController{}, "delete:Delete"),

			// get a user with id
			beego.NSRouter("/:uid", &controllers.UserController{}, "get:Get"),
		),
	)
	beego.AddNamespace(ns)
}
