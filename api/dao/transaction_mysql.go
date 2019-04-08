package dao

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type TransactionRepository interface {
	FindAll() ([]*model.Transaction, error)
	Save(user *model.Transaction)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewMysqlTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) FindAll() ([]*model.Transaction, error) {

	var transactions []*model.Transaction
	r.db.Find(&transactions)

	return transactions, r.db.Error
}

func (r *transactionRepository) Save(transaction *model.Transaction) {
	r.db.Create(&transaction)
}
