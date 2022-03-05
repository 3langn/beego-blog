package handlers

import (
	"bee-playaround1/constants"
	"bee-playaround1/helper"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func JwtFilter(ctx *context.Context) {
	ctx.Output.Header("Content-Type", "application/json")
	if strings.HasPrefix(ctx.Input.URL(), "/v1/auth") {
		return
	}

	if ctx.Input.Header("Authorization") == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		resBody, err := json.Marshal(helper.HttpException{Error: "", Message: "Authorization header is required"})
		ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}

	// Parse the token
	tokenString := ctx.Input.Header("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err != nil {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		resBody, err := json.Marshal(helper.HttpException{Error: err.Error(), Message: constants.INVALID_TOKEN_ERROR})
		ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && claims != nil {
		ctx.Input.SetData("user_id", int64(claims["user_id"].(float64)))
		return
	} else {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		resBody, err := json.Marshal(helper.HttpException{Error: constants.INVALID_TOKEN_ERROR, Message: constants.INVALID_TOKEN_ERROR})
		ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}

}
