package models

// Don't use adapter/orm => https://github.com/beego/beego/issues/4683
import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64   `orm:"auto" json:"id"`
	Username string  `orm:"unique;size(255)" json:"username"`
	Password string  `orm:"size(255)" json:"-"`
	Posts    []*Post `orm:"reverse(many)" json:"posts,omitempty"`
}

func (u *User) TableName() string {
	return "users"
}
func (u *User) Save(username string, password string) error {
	*u = User{
		Username: username,
		Password: password,
	}
	h, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 13)
	u.Password = string(h)

	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func (u *User) GetAll() ([]*User, error) {
	var users []*User
	if _, err := orm.NewOrm().QueryTable("users").All(&users); err != nil {
		panic(err)
		return nil, err
	}
	return users, nil
}

func (u *User) FindByUsername(username string) error {
	o := orm.NewOrm()
	if err := o.QueryTable("users").Filter("username", username).One(u); err != nil {
		return err
	}
	return nil
}

func (u *User) FindById(id int64) error {
	*u = User{
		Id: id,
	}
	o := orm.NewOrm()
	err := o.Read(u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) CheckPassword(password string) error {
	fmt.Println(u.Password)
	fmt.Println(password)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
