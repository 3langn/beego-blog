package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

type CategoryDto struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type Category struct {
	Id    int64   `orm:"auto" json:"id"`
	Title string  `orm:"size(75)" json:"title"`
	Slug  string  `orm:"size(100)" json:"slug"`
	Posts []*Post `orm:"reverse(many)" json:"posts,omitempty"`
}

//func (c *Category) TableName() string {
//	return "categories"
//}

func (c *Category) Create(dto []CategoryDto) ([]*Category, error) {
	o := orm.NewOrm()
	var categories []*Category
	for _, v := range dto {
		categories = append(categories, &Category{
			Title: v.Title,
			Slug:  v.Slug,
		})
	}

	_, err := o.InsertMulti(1, categories)
	if err != nil {
		return nil, err
	}
	fmt.Println(categories[0])
	return categories, nil
}

// get category by id
func (c *Category) GetById(id int64) error {
	if err := orm.NewOrm().QueryTable("category").Filter("id", id).One(c); err != nil {
		return err
	}
	return nil
}

func (c *Category) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Category) Delete() error {
	if _, err := orm.NewOrm().Delete(c); err != nil {
		return err
	}
	return nil
}

func (c Category) GetAll() ([]Category, error) {
	o := orm.NewOrm()
	var categories []Category
	_, err := o.QueryTable("category").All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *Category) GetAllCategory(order string, offset int64, limit int64) ([]Category, error) {
	var category []Category

	o := orm.NewOrm()
	qs := o.QueryTable("category")

	_, err := qs.OrderBy(order).Limit(limit, offset).All(&category)
	if err != nil {
		return nil, err
	}
	return category, nil
}
