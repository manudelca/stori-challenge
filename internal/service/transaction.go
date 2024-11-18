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
	log.Printf("Processing account info %+v", accountInfo)
	if accountInfo == nil {
		log.Printf("Account info not found")
		accountInfo = &domain.AccountInfo{
			TotalBalance: 0,
		}
	}
	monthInfo := t.accountInfoRepo.GetMonthInfo(transaction.Month)
	if monthInfo == nil {
		monthInfo = &domain.MonthInfo{
			Month:                    transaction.Month,
			TotalDebitTransactions:   0,
			TotalCreditTransactions:  0,
			NumberDebitTransactions:  0,
			NumberCreditTransactions: 0,
		}
	}
	switch transaction.MethodType {
	case domain.MethodTypeDebit:
		accountInfo.TotalBalance -= transaction.Amount
		monthInfo.TotalDebitTransactions += transaction.Amount
		monthInfo.NumberDebitTransactions += 1
	case domain.MethodTypeCredit:
		accountInfo.TotalBalance += transaction.Amount
		monthInfo.TotalCreditTransactions += transaction.Amount
		monthInfo.NumberCreditTransactions += 1
	}
	log.Printf("AccountInfo: %+v", accountInfo)
	log.Printf("MonthInfo: %+v", monthInfo)
	t.transactionRepo.SaveTransaction(transaction)
	t.accountInfoRepo.SaveAccountInfo(*accountInfo)
	t.accountInfoRepo.SaveMonthInfo(*monthInfo)
	log.Printf("Success on transaction %+v", transaction)
}
