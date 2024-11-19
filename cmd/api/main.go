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

	// Repositories
	report := repository.NewReportRepository()
	accountInfoRepo := repository.NewAccountInfoRepository()
	transactionRepo := repository.NewTransactionRepository()

	// Services
	transactionSvc := service.NewTransactionService(transactionRepo, accountInfoRepo)
	reportSvc := service.NewReportService(report, accountInfoRepo)

	// Controllers
	transactionCtrl := controller.NewTransactionController(transactionSvc)
	reportConsumer := controller.NewReportConsumerController(reportSvc)

	// ------- Este código viene a representar la ejecución de un Stream ---------
	// La idea es que se suba el archivo a la nube y esto dispare un llamado contra este servicio
	// cada cierta cantidad de bytes para ir procesando poco a poco el archivo (por simplicidad, acá se hace por línea)
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("error opening file: %v", err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	for {
		record, readErr := csvReader.Read()
		if readErr != nil {
			if readErr.Error() == "EOF" {
				break
			}
			log.Print(readErr)
		}
		// Este es el llamado que se haría via API/Stream
		transactionCtrl.ProcessRecord(record)
	}
	// --------------------------------------------------------------

	// Este llamado sería disparado una vez termina el procesamiento del archivo via cola asíncrona
	reportConsumer.SendReport()
	log.Printf("finish")
}
