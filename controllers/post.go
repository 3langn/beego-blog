package controllers

import (
	"bee-playaround1/constants"
	"bee-playaround1/dtos"
	"bee-playaround1/helper"
	"bee-playaround1/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type PostController struct {
	beego.Controller
}

// @Title CreatePost
// @Description create post
// @Success 200 {int} body models.User
// @Param	body body dtos.CreatePostDto	true	"body for post"
// @Failure 500 Internal Server Error
// @router / [post]
func (p *PostController) CreatePost() {
	var createPostDto dtos.CreatePostDto
	var post models.Post
	var user models.User
	json.Unmarshal(p.Ctx.Input.RequestBody, &createPostDto)
	err := user.FindById(int64(createPostDto.UserId))
	if err != nil {
		helper.NewHttpException(&p.Controller, constants.CREATE_POST_ERROR, err, http.StatusInternalServerError)
		return
	}

	err = post.CreatePost(createPostDto.Content, createPostDto.Title, &user)
	if err != nil {
		helper.NewHttpException(&p.Controller, constants.CREATE_POST_ERROR, err, http.StatusInternalServerError)
		return
	}

	p.Data["json"] = post
	p.ServeJSON()
}

// @Title GetPosts
// @Description create post
// @Success 200 {int} body models.Post
// @Failure 500 Internal Server Error
// @router / [get]
func (p *PostController) GetPosts() {
	var post models.Post
	posts, err := post.GetPosts()
	if err != nil {
		helper.NewHttpException(&p.Controller, constants.GET_POST_ERROR, err, http.StatusInternalServerError)
		return
	}
	p.Data["json"] = posts
	p.ServeJSON()
}

// @Title GetPostById
// @Description get Post by id
// @Param	id		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Post
// @Failure 500 Internal Server Error
// @router /:id [get]
func (p *PostController) GetPost(id int64) {
	var postModel models.Post
	post, err := postModel.GetPost(id)
	if err != nil {
		helper.NewHttpException(&p.Controller, constants.GET_POST_ERROR, err, http.StatusInternalServerError)
		return
	}
	p.Data["json"] = post
	p.ServeJSON()
}

// @Title DeletePost
// @Description delete Post by id
// @Param	id		path 	int64	true		"The key for staticblock"
// @Success 200 {string} delete success!
// @Failure 500 Internal Server Error
// @router /:id [delete]
func (p *PostController) DeletePost(id int64) {
	var postModel models.Post
	err := postModel.DeletePost(id)
	if err != nil {
		helper.NewHttpException(&p.Controller, constants.DELETE_POST_ERROR, err, http.StatusInternalServerError)
		return
	}
	p.Data["json"] = map[string]string{"Message": constants.DELETE_POST_SUCCESS}
	p.ServeJSON()
}
