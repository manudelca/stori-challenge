package repository

import "github.com/manudelca/stori-challenge/internal/domain"

type AccountInfoRepository interface {
	GetAccountInfo() *domain.AccountInfo
	SaveAccountInfo(accountInfo domain.AccountInfo)
	UpdateAccountInfo(accountInfo domain.AccountInfo)
	GetMonthInfo(month int) *domain.MonthInfo
	SaveMonthInfo(monthInfo domain.MonthInfo)
	UpdateMonthInfo(monthInfo domain.MonthInfo)
}

type accountInfoRepository struct {
	accountInfoStorage *domain.AccountInfo
	MonthInfo          map[int]domain.MonthInfo
}

func NewAccountInfoRepository() AccountInfoRepository {
	return &accountInfoRepository{
		accountInfoStorage: &domain.AccountInfo{},
		MonthInfo:          make(map[int]domain.MonthInfo, 12),
	}
}

func (a *accountInfoRepository) SaveAccountInfo(accountInfo domain.AccountInfo) {
	a.accountInfoStorage = &accountInfo
}

func (a *accountInfoRepository) UpdateAccountInfo(accountInfo domain.AccountInfo) {
	a.accountInfoStorage = &accountInfo
}

func (a *accountInfoRepository) GetAccountInfo() *domain.AccountInfo {
	return a.accountInfoStorage
}

func (a *accountInfoRepository) GetMonthInfo(month int) *domain.MonthInfo {
	MonthInfo := a.MonthInfo[month]
	return &MonthInfo
}

func (a *accountInfoRepository) SaveMonthInfo(monthInfo domain.MonthInfo) {
	a.MonthInfo[monthInfo.Month] = monthInfo
}

func (a *accountInfoRepository) UpdateMonthInfo(monthInfo domain.MonthInfo) {
	a.MonthInfo[monthInfo.Month] = monthInfo
}
