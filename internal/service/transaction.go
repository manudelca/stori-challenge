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
	transactionRepo repository.TransactionRepository
	accountInfoRepo repository.AccountInfoRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository, accountInfoRepo repository.AccountInfoRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		accountInfoRepo: accountInfoRepo,
	}
}

func (t *transactionService) ProcessTransaction(transaction domain.Transaction) {
	log.Printf("Processing transaction %+v", transaction)
	accountInfo := t.accountInfoRepo.GetAccountInfo()
	if accountInfo == nil {
		accountInfo = &domain.AccountInfo{}
	}
	monthInfo := t.accountInfoRepo.GetMonthInfo(transaction.Month)
	if monthInfo == nil {
		monthInfo = &domain.MonthInfo{}
	}
	switch transaction.MethodType {
	case domain.MethodTypeDebit:
		accountInfo.TotalBalance -= transaction.Amount
	case domain.MethodTypeCredit:
		accountInfo.TotalBalance += transaction.Amount
	}
	accountInfo.TotalBalance += transaction.Amount
	t.transactionRepo.SaveTransaction(transaction)
	t.accountInfoRepo.UpdateAccountInfo(*accountInfo)
	t.accountInfoRepo.UpdateMonthInfo(*monthInfo)
	log.Printf("Success on transaction %+v", transaction)
}
