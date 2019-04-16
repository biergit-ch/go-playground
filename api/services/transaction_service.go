package services

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/repo"
	"git.skydevelopment.ch/zrh-dev/go-basics/models"
	log "github.com/sirupsen/logrus"
)

type TransactionService interface {
	GetAllTransactions() []*models.Transaction
	GetTransaction(id int) *models.Transaction
	CreateTransaction(user *models.Transaction)
}

type transactionService struct {
	repo repo.TransactionRepository
}

func NewTransactionService(r repo.TransactionRepository) TransactionService {
	return &transactionService{
		repo: r,
	}
}

func (s *transactionService) GetAllTransactions() []*models.Transaction {
	log.Debug("Try to query the repo for the transactions" )
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
