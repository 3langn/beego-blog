package models

import "github.com/beego/beego/v2/client/orm"

type Category struct {
	Id    int64   `orm:"auto" json:"id"`
	Title string  `orm:"size(75)" json:"title"`
	Slug  string  `orm:"size(100)" json:"slug"`
	Posts []*Post `json:"posts"`
}

// create table category
func (c *Category) TableName() string {
	return "categories"
}

func (c *Category) Create() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

// get category by id
func (c *Category) GetById(id int64) error {
	if err := orm.NewOrm().QueryTable(c.TableName()).Filter("id", id).One(c); err != nil {
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
	_, err := o.QueryTable(c.TableName()).All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *Category) GetAllCategory(order string, offset int64, limit int64) ([]Category, error) {
	var category []Category

	o := orm.NewOrm()
	qs := o.QueryTable(c.TableName())

	_, err := qs.OrderBy(order).Limit(limit, offset).All(&category)
	if err != nil {
		return nil, err
	}
	return category, nil
}
