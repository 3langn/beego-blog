package dtos

import "bee-playaround1/models"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string      `json:"message"`
	User    models.User `json:"user"`
}
