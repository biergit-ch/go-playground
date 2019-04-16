package services

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/dao"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	log "github.com/sirupsen/logrus"
)

type TransactionService interface {
	GetAllTransactions() []*models.Transaction
	GetTransaction(id int) *models.Transaction
	CreateTransaction(user *models.Transaction)
}

type transactionService struct {
	repo dao.TransactionRepository
}

func NewTransactionService(r dao.TransactionRepository) TransactionService {
	return &transactionService{
		repo: r,
	}
}

func (s *transactionService) GetAllTransactions() []*models.Transaction {

	transactions, err := s.repo.FindAll()
	if err != nil {
		log.Fatal("Failed to get all users from repo")
	}
	return transactions
}

func (s *transactionService) GetTransaction(id int) *models.Transaction {

	transaction, err := s.repo.FindOne(id)
	if err != nil {
		log.Fatal("Failed to get all users from repo")
	}
	return transaction
}

func (s *transactionService) CreateTransaction(t *models.Transaction) {
	s.repo.Save(t)
}
