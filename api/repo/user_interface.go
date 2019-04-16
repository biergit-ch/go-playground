package repo

import "git.skydevelopment.ch/zrh-dev/go-basics/models"

type UserRepository interface {
	FindAll() ([]*models.User, error)
	Save(user *models.User) *models.User
	Update(user *models.User) (*models.User, error)
	Delete(id int) error
	FindOne(id string) (*models.User, error)
}
