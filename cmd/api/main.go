package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/manudelca/stori-challenge/internal/controller"
	"github.com/manudelca/stori-challenge/internal/repository"
	"github.com/manudelca/stori-challenge/internal/service"
)

func main() {

	// Repositories
	//report := repository.NewReportRepository()
	accountInfoRepo := repository.NewAccountInfoRepository()
	transactionRepo := repository.NewTransactionRepository()

	// Services
	transactionSvc := service.NewTransactionService(transactionRepo, accountInfoRepo)
	//reportSvc := service.NewReportService(report, accountInfoRepo)

	// Controllers
	transactionCtrl := controller.NewTransactionController(transactionSvc)
	//reportConsumer := controller.NewReportConsumerController(reportSvc)

	// Lo que hace AWS Stream (?
	/*
		fileName := os.Args[1]
		file, err := os.Open(fileName)
		if err != nil {
			log.Printf("error opening file: %v", err)
			return
		}
		defer file.Close()

		csvReader := csv.NewReader(file)
		for {
			record, readErr := csvReader.Read() // Read one record at a time
			if readErr != nil {
				if readErr.Error() == "EOF" {
					break // End of file
				}
				log.Print(readErr) // Some other error
			}
			// Process the record (each record is a slice of strings)
			transactionCtrl.ProcessRecord(record)
		}*/

	//reportConsumer.SendReport()
	lambda.Start(transactionCtrl.Handler)
	log.Printf("finish")
}
