package controllers

import (
	"bee-playaround1/dtos"
	"bee-playaround1/models"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body body dtos.Login	true	"body for user content"
// @Success 200 {int} body models.User
// @Failure 403 body is empty
// @router /register [post]
func (u *UserController) Register() {
	var user models.User
	var login dtos.Login
	json.Unmarshal(u.Ctx.Input.RequestBody, &login)
	err := user.Save(login.Username, login.Password)
	if err != nil {
		u.Data["json"] = map[string]string{"message": "User name already exists"}
	} else {
		u.Data["json"] = map[string]models.User{"user": user}
	}
	u.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body body dtos.Login	true	"body for user content"
// @Success 200 {int} body models.User
// @Failure 400 Login failed
// @router /login [post]
func (u *UserController) Login() {
	var login dtos.Login
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &login)
	err := user.FindByUsername(login.Username)
	if err != nil {
		fmt.Println(err, "err")
		models.NewInternalException(&u.Controller, "Khong tim thay ten tai khoan", err)
		return
	}
	err = user.CheckPassword(login.Password)
	if err != nil {
		models.NewInternalException(&u.Controller, "Sai mai khau", err)
		return
	}
	u.Data["json"] = &dtos.LoginResponse{Message: "Dang nhap thanh cong", User: user}
	u.ServeJSON()
}

// @Title Get all users
// @Description Get all users
// @Success 200 {int} body []*models.User
// @Failure 403 body is empty
// @router / [get]
func (u *UserController) GetAll() {
	var users models.User
	all, err := users.GetAll()
	if err != nil {
		u.Data["json"] = map[string]string{"message": err.Error()}
	} else {
		u.Data["json"] = map[string][]*models.User{"users": all}
	}
	u.ServeJSON()
}
