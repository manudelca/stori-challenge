package controller

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"strconv"
	"time"

	"github.com/manudelca/stori-challenge/internal/domain"
	"github.com/manudelca/stori-challenge/internal/service"
)

type TransactionController interface {
	//S3EventHandler(ctx context.Context, s3Event events.S3Event)
	Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
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

func (t *transactionController) Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "\"Hello from Lambda!\"",
	}
	return response, nil
}

/*
	func (t *transactionController) S3EventHandler(ctx context.Context, s3Event events.S3Event) {
		for _, record := range s3Event.Records {
			s3 := record.S3
			fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		}
	}
*/
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
