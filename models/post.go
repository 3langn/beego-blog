package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type PostDto struct {
	Title      string        `json:"title"`
	Content    string        `json:"content"`
	Categories []CategoryDto `json:"categories"`
}

type Post struct {
	Id         int64       `orm:"auto" json:"id"`
	Title      string      `orm:"size(100)" json:"title"`
	Content    string      `orm:"size(100)" json:"content"`
	Author     *User       `orm:"rel(fk);null;on_delete(set_null)" json:"author"`
	Categories []*Category `orm:"rel(m2m)" json:"categories,omitempty"`
	Created    time.Time   `orm:"auto_now_add;type(datetime)" json:"created_at"`
	Updated    time.Time   `orm:"auto_now;type(datetime)" json:"updated_at"`
}

func (p *Post) TableName() string {
	return "posts"
}

func (p *Post) CreatePost(title string, content string, user *User, categories []*Category) error {

	*p = Post{
		Title:      title,
		Content:    content,
		Author:     user,
		Categories: categories,
	}
	p.Author = user
	o := orm.NewOrm()
	_, err := o.Insert(p)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func (p *Post) GetPosts(categoryTitle string, offset int, limit int) ([]*Post, error) {
	var posts []*Post
	o := orm.NewOrm()
	qs := o.QueryTable("posts")
	if categoryTitle != "" {
		qs = qs.Filter("Categories__Category__Title", categoryTitle)
	}
	_, err := qs.Offset(offset).Limit(limit).RelatedSel().All(&posts)

	for _, post := range posts {
		_, err := o.LoadRelated(post, "Categories")
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *Post) GetPost(id int64) (*Post, error) {
	var post Post
	o := orm.NewOrm()
	err := o.QueryTable(p.TableName()).RelatedSel("author", "category").Filter("id", id).One(&post)
	if err != nil {
		fmt.Println(err, ":GetPostById")
		return nil, err
	}
	return &post, nil
}

// update post
func (p *Post) Update(id int64) error {
	o := orm.NewOrm()
	p.Id = id
	if o.Read(p) == nil {
		if _, err := o.Update(&p); err != nil {
			return err
		}
	}
	return nil
}

func (p *Post) DeletePost(id int64, userId int64) error {
	*p = Post{Id: id}
	o := orm.NewOrm()
	rs, err := o.Raw("DELETE FROM posts WHERE posts.id = ? AND posts.author_id = ?", id, userId).Exec()
	rowsAffected, err := rs.RowsAffected()
	fmt.Println(rowsAffected, err)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("can't delete post")
	}
	return nil
}

func (p *Post) FindPostById(id int64) (*Post, error) {
	post := &Post{
		Id: id,
	}
	o := orm.NewOrm()
	err := o.Read(post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
