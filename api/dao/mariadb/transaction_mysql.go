package mariadb

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/repo"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewMysqlTransactionRepository(db *gorm.DB) repo.TransactionRepository {
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
