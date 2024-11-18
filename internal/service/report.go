package service

import (
	"github.com/manudelca/stori-challenge/internal/domain"
	"github.com/manudelca/stori-challenge/internal/repository"
	"log"
	"sort"
)

type ReportService interface {
	SendReport()
}

type reportService struct {
	reportRepo      repository.ReportRepository
	accountInfoRepo repository.AccountInfoRepository
}

func NewReportService(reportRepo repository.ReportRepository, accountInfoRepo repository.AccountInfoRepository) ReportService {
	return &reportService{
		reportRepo:      reportRepo,
		accountInfoRepo: accountInfoRepo,
	}
}

func (r *reportService) SendReport() {
	log.Printf("Processing report")
	accountInfo := r.accountInfoRepo.GetAccountInfo()
	if accountInfo == nil {
		log.Printf("No account info found")
		return
	}
	months := r.accountInfoRepo.SearchAllMonthInYear()
	sort.Sort(domain.ByMonth(months))
	r.reportRepo.SendReport(*accountInfo, months)
	log.Printf("report process success")
}
