package dao

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

type UserRepository interface {
	FindAll() ([]*models.User, error)
	Save(user *models.User) *models.User
	Update(user *models.User) (*models.User, error)
	Delete(id int64) error
	Find(id int64) ([]*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindAll() ([]*models.User, error) {

	var users []*models.User
	r.db.Find(&users)

	return users, r.db.Error
}

func (r *userRepository) Find(id int64) ( []*models.User, error) {
	log.Debug("Find user with id ", id, " in mysql database")
	var users []*models.User

	r.db.First(&users, id)

	return users, r.db.Error
}

func (r *userRepository) Save(user *models.User) *models.User {
	r.db.Create(&user)
	return user
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {

	r.db.Update(&user)

	return user, r.db.Error
}

func (r *userRepository) Delete(id int64) error {

	log.WithFields(log.Fields{"user_id": id}).Debug("delete user from mysql db")

	r.db.Where("id = ?", id).Delete(&models.User{})

	return r.db.Error
}
