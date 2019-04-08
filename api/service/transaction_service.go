package service

import (
	"git.skydevelopment.ch/zrh-dev/go-basics/api/dao"
	"git.skydevelopment.ch/zrh-dev/go-basics/api/model"
	log "github.com/sirupsen/logrus"
)

type TransactionService interface {
	GetAllTransactions() []*model.Transaction
	CreateTransaction(user *model.Transaction)
}

type transactionService struct {
	repo dao.TransactionRepository
}

func NewTransactionService(r dao.TransactionRepository) TransactionService {
	return &transactionService{
		repo: r,
	}
}

func (s *transactionService) GetAllTransactions() []*model.Transaction {

	transactions, err := s.repo.FindAll()
	if err != nil {
		log.Fatal("Failed to get all users from repo")
	}
	return transactions
}

func (s *transactionService) CreateTransaction(t *model.Transaction) {
	s.repo.Save(t)
}
