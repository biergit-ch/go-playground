package dao

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

type UserRepository interface {
	FindAll() ([]*model.User, error)
	Save(user *model.User)
}

type userRepository struct {
	db *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindAll() ([]*model.User, error) {

	var users []*model.User
	r.db.Find(&users)

	return users, r.db.Error
}

func (r *userRepository) Save(user *model.User) {
	r.db.Create(&user)
}
