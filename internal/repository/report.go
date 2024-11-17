package repository

type ReportRepository interface {
}

type reportRepository struct {
}

func NewReportRepository() ReportRepository {
	return &reportRepository{}
}
