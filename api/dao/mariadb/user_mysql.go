package mariadb

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/repo"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type userRepository struct {
	db *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) repo.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindAll() ([]*models.User, error) {
	var users []*models.User
	r.db.Find(&users)
	return users, r.db.Error
}

func (r *userRepository) FindOne(id string) (*models.User, error) {
	log.Debug("Find user with id ", id, " in mysql database")

	// convert string to int
	var userId, _ = strconv.Atoi(id)
	var users []*models.User
	r.db.First(&users, userId)

	if len(users) > 0 {
		return users[0], nil
	} else {
		return nil, r.db.Error
	}
}

func (r *userRepository) Save(user *models.User) *models.User {
	r.db.Create(&user)
	return user
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {
	r.db.Update(&user)
	return user, r.db.Error
}

func (r *userRepository) Delete(id int) error {
	log.WithFields(log.Fields{"user_id": id}).Debug("delete user from mysql db")
	r.db.Where("id = ?", id).Delete(&models.User{})
	return r.db.Error
}
