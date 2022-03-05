// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// @Security ApiKeyAuth
// @SecurityDefinition ApiKeyAuth apiKey Authorization header "Authorization token"
package routers

import (
	"bee-playaround1/controllers"
	"bee-playaround1/handlers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// Swagger API only support NSNamespace

	beego.InsertFilter("/*", beego.BeforeRouter, handlers.JwtFilter)
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/auth",
			beego.NSInclude(&controllers.AuthController{}),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(&controllers.UserController{}),
		),
		beego.NSNamespace("/post",
			beego.NSInclude(&controllers.PostController{}),
		),
		beego.NSNamespace("/categories",
			beego.NSInclude(&controllers.CategoryController{}),
		),
	)

	beego.AddNamespace(ns)
}
