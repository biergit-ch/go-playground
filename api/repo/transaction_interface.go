package repo

import "git.skydevelopment.ch/zrh-dev/go-basics/models"

type TransactionRepository interface {
	FindAll() ([]*models.Transaction, error)
	FindOne(id int) (*models.Transaction, error)
	Save(user *models.Transaction)
}
