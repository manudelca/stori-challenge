package repository

import (
	"github.com/manudelca/stori-challenge/internal/domain"
)

type AccountInfoRepository interface {
	GetAccountInfo() *domain.AccountInfo
	SaveAccountInfo(accountInfo domain.AccountInfo)
	GetMonthInfo(month int) *domain.MonthInfo
	SaveMonthInfo(monthInfo domain.MonthInfo)
	SearchAllMonthInYear() []domain.MonthInfo
}

type accountInfoRepository struct {
	// Esto viene a representar 2 storage por separado de tipo NoSQL para mayor escalabilidad
	// (1 para el accountInfo otro para el almacenamiento de meses)
	accountInfoStorage *domain.AccountInfo
	MonthInfo          map[int]domain.MonthInfo
}

func NewAccountInfoRepository() AccountInfoRepository {
	return &accountInfoRepository{
		MonthInfo: make(map[int]domain.MonthInfo, 12),
	}
}

func (a *accountInfoRepository) SaveAccountInfo(accountInfo domain.AccountInfo) {
	a.accountInfoStorage = &accountInfo
}

func (a *accountInfoRepository) GetAccountInfo() *domain.AccountInfo {
	return a.accountInfoStorage
}

func (a *accountInfoRepository) GetMonthInfo(month int) *domain.MonthInfo {
	MonthInfo, exists := a.MonthInfo[month]
	if !exists {
		return nil
	}
	return &MonthInfo
}

func (a *accountInfoRepository) SaveMonthInfo(monthInfo domain.MonthInfo) {
	a.MonthInfo[monthInfo.Month] = monthInfo
}

func (a *accountInfoRepository) SearchAllMonthInYear() []domain.MonthInfo {
	response := make([]domain.MonthInfo, 0, 12)
	for _, value := range a.MonthInfo {
		response = append(response, value)
	}
	return response
}
