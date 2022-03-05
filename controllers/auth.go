package controllers

import (
	"bee-playaround1/constants"
	"bee-playaround1/dtos"
	"bee-playaround1/helper"
	"bee-playaround1/models"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type AuthController struct {
	beego.Controller
}

// @Title Register
// @Description Register users
// @Param	body body dtos.Login	true	"body for user content"
// @Success 200 {int} body models.User
// @Failure 403 body is empty
// @router /register [post]
func (a *AuthController) Register() {
	var user models.User
	var login dtos.Login
	json.Unmarshal(a.Ctx.Input.RequestBody, &login)
	err := user.Save(login.Username, login.Password)
	if err != nil {
		helper.NewHttpException(&a.Controller, constants.DUPLICATE_EMAIL_ERROR, err, http.StatusBadRequest)
		return
	} else {
		a.Data["json"] = map[string]models.User{"user": user}
	}
	a.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	body body dtos.Login	true	"body for user content"
// @Success 200 {int} body models.User
// @Failure 400 Login failed
// @router /login [post]
func (a *AuthController) Login() {
	var login dtos.Login
	var user models.User
	json.Unmarshal(a.Ctx.Input.RequestBody, &login)
	err := user.FindByUsername(login.Username)
	if err != nil {
		fmt.Println(err, "err")
		helper.NewHttpException(&a.Controller, constants.EMAIL_NOT_FOUND_ERROR, err, http.StatusInternalServerError)
		return
	}

	err = user.CheckPassword(login.Password)
	if err != nil {
		helper.NewHttpException(&a.Controller, constants.INVALID_PASSWORD_ERROR, err, http.StatusUnauthorized)
		return
	}

	a.Data["json"] = &dtos.LoginResponse{Message: constants.LOGIN_SUCCESS, User: user}
	a.ServeJSON()
}
