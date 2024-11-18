package controller

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/manudelca/stori-challenge/internal/domain"
	"github.com/manudelca/stori-challenge/internal/service"
)

type TransactionController interface {
	ProcessRecord(record []string)
}

type transactionController struct {
	svc service.TransactionService
}

func NewTransactionController(svc service.TransactionService) TransactionController {
	return &transactionController{
		svc: svc,
	}
}

func (t *transactionController) ProcessRecord(record []string) {
	// TODO: Skip first record
	transaction, buildErr := buildTransactionFromRecord(record)
	if buildErr != nil {
		log.Printf("failed to build transaction from record: %v", buildErr)
		return
	}
	t.svc.ProcessTransaction(*transaction)
}

func buildTransactionFromRecord(record []string) (*domain.Transaction, error) {
	// TODO: Add validations

	methodType, ok := map[string]domain.MethodType{
		"-": domain.MethodTypeDebit,
		"+": domain.MethodTypeCredit,
	}[record[2][0:1]]
	if !ok {
		return nil, fmt.Errorf("failed to parse method from: %s", record[2][0:1])
	}

	strAmount := record[2][1:]
	amount, err := strconv.ParseFloat(strAmount, 64)
	if err != nil {
		return nil, err
	}
	date, dateErr := parseDate(record[1])
	if dateErr != nil {
		return nil, dateErr
	}

	return &domain.Transaction{
		ID:         record[0],
		Day:        date.Day(),
		Month:      int(date.Month()),
		MethodType: methodType,
		Amount:     amount,
	}, nil
}

func parseDate(input string) (time.Time, error) {
	layout := "1/2"
	parsedDate, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, err
	}

	return parsedDate, nil
}
