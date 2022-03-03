package models

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type HttpException struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewInternalException(c *beego.Controller, message string, error error) {
	c.Ctx.Output.SetStatus(http.StatusBadRequest)
	c.Data["json"] = &HttpException{message, error.Error()}
	c.ServeJSON()
}
