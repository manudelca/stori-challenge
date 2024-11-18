package service

import (
	"github.com/manudelca/stori-challenge/internal/repository"
	"log"
)

type ReportService interface {
	SendReport()
}

type reportService struct {
	reportRepo repository.ReportRepository
}

func NewReportService(reportRepo repository.ReportRepository) ReportService {
	return &reportService{
		reportRepo: reportRepo,
	}
}

func (r *reportService) SendReport() {
	log.Printf("Processing report")
}
