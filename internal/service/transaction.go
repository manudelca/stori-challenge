package service

import (
	"log"

	"github.com/manudelca/stori-challenge/internal/domain"
	"github.com/manudelca/stori-challenge/internal/repository"
)

type TransactionService interface {
	ProcessTransaction(transaction domain.Transaction)
}

type transactionService struct {
	reportRepo repository.ReportRepository
}

func (t transactionService) ProcessTransaction(transaction domain.Transaction) {
	log.Printf("Processing transaction %+v", transaction)
}

func NewTransactionService(reportRepo repository.ReportRepository) TransactionService {
	return &transactionService{
		reportRepo: reportRepo,
	}
}
