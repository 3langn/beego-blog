package controllers

import (
	"bee-playaround1/constants"
	"bee-playaround1/helper"
	"bee-playaround1/models"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type PostController struct {
	beego.Controller
}

// @Title CreatePost
// @Description create post
// @Success 200 {int} body models.User
// @Param	body body models.PostDto	true	"body for post"
// @Failure 500 Internal Server Error
// @router / [post]
func (p *PostController) CreatePost() {
	var createPostDto models.PostDto
	var post models.Post
	var user models.User
	var category models.Category
	json.Unmarshal(p.Ctx.Input.RequestBody, &createPostDto)

	userId := p.Ctx.Input.GetData("user_id")

	err := user.FindById(userId.(int64))
	if err != nil {
		helper.NewHttpException(&p.Controller, constants.CREATE_POST_ERROR, err, http.StatusInternalServerError)
		return
	}

	categories, err := category.Create(createPostDto.Categories)
	if err != nil {
		helper.NewHttpException(&p.Controller, constants.CREATE_POST_ERROR, err, http.StatusInternalServerError)
		return
	}

	err = post.CreatePost(createPostDto.Content, createPostDto.Title, &user, categories)
	if err != nil {
		helper.NewHttpException(&p.Controller, constants.CREATE_POST_ERROR, err, http.StatusInternalServerError)
		return
	}

	o := orm.NewOrm()
	_, err = o.QueryM2M(&post, "Categories").Add(categories)
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
// @Param 	offset 	query 	int 	false 	"offset"
// @Param 	limit 	query 	int 	false 	"limit"
// @Param   category_title 	query 	string 	false 	"category_title"
// @Failure 500 Internal Server Error
// @router / [get]
func (p *PostController) GetPosts() {
	var post models.Post
	categoryTitle := p.GetString("category_title")
	offset, _ := p.GetInt("offset")
	limit, _ := p.GetInt("limit")
	posts, err := post.GetPosts(categoryTitle, offset, limit)
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
// @Failure 404 not found resource
// @router /:id [get]
func (p *PostController) GetPost(id int64) {
	var postModel models.Post
	post, err := postModel.GetPost(id)
	if err != nil {
		helper.NewHttpException(&p.Controller, constants.GET_POST_ERROR, err, http.StatusNotFound)
		return
	}
	p.Data["json"] = post
	p.ServeJSON()
}

// @Title UpdatePost
// @Description update the post
// @Param	id		path 	int64	true		"The id you want to update"
// @Param	body		body 	models.PostDto	true		"body for post content"
// @Success 200 {object} models.Post
// @Failure 403 :id is not int
// @router /:id [put]
func (p *PostController) UpdatePost(id int64) {
	var post models.Post
	json.Unmarshal(p.Ctx.Input.RequestBody, &post)
	err := post.Update(id)
	if err != nil {
		helper.NewHttpException(&p.Controller, constants.UPDATE_POST_ERROR, err, http.StatusInternalServerError)
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
	userId := p.Ctx.Input.GetData("user_id")

	err := postModel.DeletePost(id, userId.(int64))

	if err != nil {
		helper.NewHttpException(&p.Controller, constants.DELETE_POST_ERROR, err, http.StatusBadRequest)
		return
	}
	p.Data["json"] = map[string]string{"Message": constants.DELETE_POST_SUCCESS}
	p.ServeJSON()
}
