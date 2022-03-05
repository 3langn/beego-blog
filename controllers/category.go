package controllers

import (
	"bee-playaround1/constants"
	"bee-playaround1/helper"
	"bee-playaround1/models"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"strconv"
)

type CategoryController struct {
	beego.Controller
}

// CRUD
// @Title Create
// @Description create category
// @Param	body		body 	models.Category	true		"body for category content"
// @Success 200 {int} models.Category.Id
// @Failure 500 internal server error
// @router / [post]
func (c *CategoryController) Create() {
	var category models.Category
	if err := c.ParseForm(&category); err != nil {
		helper.NewHttpException(&c.Controller, "Can't create category", err, http.StatusBadRequest)
	} else {
		c.Data["json"] = category.Create()
	}
	c.ServeJSON()
}

// @Title Get
// @Description get category by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Category
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CategoryController) Get() {
	errMessage := "Can't get category"
	id := c.Ctx.Input.Param(":id")
	if id != "" {
		helper.NewHttpException(&c.Controller, errMessage, nil, http.StatusBadRequest)
		return
	}

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		helper.NewHttpException(&c.Controller, errMessage, nil, http.StatusBadRequest)
		return
	}

	var category models.Category
	err = category.GetById(i)
	if err != nil {
		helper.NewHttpException(&c.Controller, "Category not found", err, http.StatusNotFound)
		return
	}
	c.ServeJSON()
}

// @Title GetAll
// @Description get category
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Category
// @Failure 500 Internal Server Error
// @router / [get]
func (c *CategoryController) GetAllCategory() {
	var category models.Category
	var order string
	var limit int64 = 10
	var offset int64

	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	if v := c.GetString("order"); v != "" {
		order = v
	}

	l, err := category.GetAllCategory(order, offset, limit)
	if err != nil {
		helper.NewHttpException(&c.Controller, constants.GET_CATEGORIES_ERROR, err, http.StatusNotFound)
		return
	}
	c.Data["json"] = l

	c.ServeJSON()
}

//func (c CategoryController) GetAll() {
//	var category models.Category
//
//	categories, err := category.GetAll()
//	if err != nil {
//		helper.NewHttpException(&c.Controller, "Cannot get all categories", err, http.StatusInternalServerError)
//	}
//	c.Data["json"] = categories
//	c.ServeJSON()
//}
