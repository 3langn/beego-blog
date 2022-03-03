package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["bee-playaround1/controllers:PostController"] = append(beego.GlobalControllerRouter["bee-playaround1/controllers:PostController"],
        beego.ControllerComments{
            Method: "CreatePost",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-playaround1/controllers:PostController"] = append(beego.GlobalControllerRouter["bee-playaround1/controllers:PostController"],
        beego.ControllerComments{
            Method: "GetPosts",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-playaround1/controllers:PostController"] = append(beego.GlobalControllerRouter["bee-playaround1/controllers:PostController"],
        beego.ControllerComments{
            Method: "GetPost",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("id", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-playaround1/controllers:PostController"] = append(beego.GlobalControllerRouter["bee-playaround1/controllers:PostController"],
        beego.ControllerComments{
            Method: "DeletePost",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("id", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-playaround1/controllers:UserController"] = append(beego.GlobalControllerRouter["bee-playaround1/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-playaround1/controllers:UserController"] = append(beego.GlobalControllerRouter["bee-playaround1/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bee-playaround1/controllers:UserController"] = append(beego.GlobalControllerRouter["bee-playaround1/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/register",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
