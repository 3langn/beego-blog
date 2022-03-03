package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

type Post struct {
	Id      int64  `orm:"auto" json:"id"`
	Title   string `orm:"size(100)" json:"title"`
	Content string `orm:"size(100)" json:"content"`
	Author  *User  `orm:"rel(fk);null;on_delete(set_null)" json:"author"`
}

func (p *Post) TableName() string {
	return "posts"
}

func (p *Post) CreatePost(title string, content string, user *User) error {
	*p = Post{
		Title:   title,
		Content: content,
		Author:  user,
	}
	p.Author = user
	o := orm.NewOrm()
	_, err := o.Insert(p)
	if err != nil {
		fmt.Println(err, ":CreatePost")
		return err
	}
	return nil
}

func (p *Post) GetPosts() ([]*Post, error) {
	var posts []*Post
	o := orm.NewOrm()
	_, err := o.QueryTable("posts").RelatedSel("author").All(&posts)
	// query raw join author
	//_, err := o.Raw(`SELECT * FROM posts LEFT JOIN users ON users.id = posts.author_id`).QueryRows(&posts)
	if err != nil {
		fmt.Println(err, ":GetPosts")
		return nil, err
	}

	fmt.Println(posts[0])
	return posts, nil
}

func (p *Post) GetPost(id int64) (*Post, error) {
	var post Post
	o := orm.NewOrm()
	err := o.QueryTable("posts").RelatedSel("author").Filter("id", id).One(&post)
	if err != nil {
		fmt.Println(err, ":GetPostById")
		return nil, err
	}
	return &post, nil
}

func (p *Post) DeletePost(id int64) error {
	*p = Post{Id: id}
	o := orm.NewOrm()
	rs, err := o.Raw("DELETE FROM posts WHERE posts.id = ?", id).Exec()
	rowsAffected, err := rs.RowsAffected()
	fmt.Println(rowsAffected, err)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("post with id %d not found", id)
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
