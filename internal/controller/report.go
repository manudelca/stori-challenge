package controller

import "github.com/manudelca/stori-challenge/internal/service"

type ReportConsumerController interface {
	SendReport()
}

type reportConsumerController struct {
	svc service.ReportService
}

func NewReportConsumerController(reportService service.ReportService) ReportConsumerController {
	return &reportConsumerController{
		svc: reportService,
	}
}

func (r *reportConsumerController) SendReport() {
	r.svc.SendReport()
}
