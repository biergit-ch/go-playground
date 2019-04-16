package repo

import "git.skydevelopment.ch/zrh-dev/go-basics/models"

type GroupRepository interface {
	FindAll() ([]*models.Group, error)
	Save(user *models.Group)
}
