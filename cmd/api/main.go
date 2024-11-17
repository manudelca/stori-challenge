package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/manudelca/stori-challenge/internal/controller"
	"github.com/manudelca/stori-challenge/internal/repository"
	"github.com/manudelca/stori-challenge/internal/service"
)

func main() {

	report := repository.NewReportRepository()
	svc := service.NewTransactionService(report)
	ctrl := controller.NewTransactionController(svc)

	// Lo que hace AWS Stream (?
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
		ctrl.ProcessRecord(record)
	}
	log.Printf("finish")
}
