package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

type repo struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) FindAll() ([]*User, error) {

	var users []*User
	r.db.Find(&users)

	return users, r.db.Error
}

func (r *repo) Save(user *User) {
	r.db.Create(&user)
}
