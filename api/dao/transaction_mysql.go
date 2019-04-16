package dao

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TransactionRepository interface {
	FindAll() ([]*models.Transaction, error)
	FindOne(id int) (*models.Transaction, error)
	Save(user *models.Transaction)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewMysqlTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) FindAll() ([]*models.Transaction, error) {

	var transactions []*models.Transaction
	r.db.Find(&transactions)

	return transactions, r.db.Error
}

func (r *transactionRepository) FindOne(id int) (*models.Transaction, error) {

	var transactions []*models.Transaction
	r.db.First(&transactions, id)

	if len(transactions) > 0 {
		return transactions[0], nil
	} else {
		return nil, r.db.Error
	}
}

func (r *transactionRepository) Save(transaction *models.Transaction) {
	r.db.Create(&transaction)
}
