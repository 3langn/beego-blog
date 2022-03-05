package controllers

import (
	"bee-playaround1/constants"
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
// @Param	body body models.LoginDto	true	"body for user content"
// @Success 200 {int} body models.User
// @Failure 403 body is empty
// @router /register [post]
func (a *AuthController) Register() {
	var user models.User
	var login models.LoginDto
	json.Unmarshal(a.Ctx.Input.RequestBody, &login)
	err := user.Save(login.Username, login.Password)
	if err != nil {
		panic(err)
		helper.NewHttpException(&a.Controller, constants.DUPLICATE_EMAIL_ERROR, err, http.StatusBadRequest)
		return
	} else {
		a.Data["json"] = map[string]models.User{"user": user}
	}
	a.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	body body models.LoginDto	true	"body for user "
// @Success 200 {int} body models.User
// @Failure 400 Login failed
// @router /login [post]
func (a *AuthController) Login() {
	var login models.LoginDto
	var user models.User
	json.Unmarshal(a.Ctx.Input.RequestBody, &login)

	err := user.FindByUsername(login.Username)

	if err != nil {
		if (err.Error()) == "<QuerySeter> no row found" {
			helper.NewHttpException(&a.Controller, constants.USER_NOT_FOUND_ERROR, err, http.StatusNotFound)
			return
		}
		helper.NewHttpException(&a.Controller, constants.EMAIL_NOT_FOUND_ERROR, err, http.StatusInternalServerError)
		return
	}
	err = user.CheckPassword(login.Password)
	if err != nil {
		helper.NewHttpException(&a.Controller, constants.INVALID_PASSWORD_ERROR, err, http.StatusUnauthorized)
		return
	}
	var t models.Token

	token, err := t.GenerateAuthToken(user.Id)
	if err != nil {
		helper.NewHttpException(&a.Controller, constants.INVALID_PASSWORD_ERROR, err, http.StatusUnauthorized)
		return
	}
	fmt.Println("3")
	a.Data["json"] = &models.LoginResponseDto{Message: constants.LOGIN_SUCCESS, User: &user, Token: token}
	a.ServeJSON()
}
