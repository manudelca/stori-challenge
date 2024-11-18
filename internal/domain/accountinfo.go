package domain

type AccountInfo struct {
	TotalBalance float64 `json:"total_balance"`
}

type MonthInfo struct {
	Month                    int     `json:"month"`
	TotalDebitTransactions   float64 `json:"total_debit_transactions"`
	TotalCreditTransactions  float64 `json:"total_credit_transactions"`
	NumberDebitTransactions  int     `json:"number_debit_transactions"`
	NumberCreditTransactions int     `json:"number_credit_transactions"`
}

type ByMonth []MonthInfo

func (m ByMonth) Len() int           { return len(m) }
func (m ByMonth) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m ByMonth) Less(i, j int) bool { return m[i].Month < m[j].Month }
