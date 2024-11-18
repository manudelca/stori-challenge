package domain

type AccountInfo struct {
	TotalBalance float64 `json:"total_balance"`
}

type MonthInfo struct {
	Month                    int `json:"month"`
	TotalDebitTransactions   int `json:"total_debit_transactions"`
	TotalCreditTransactions  int `json:"total_credit_transactions"`
	NumberDebitTransactions  int `json:"number_debit_transactions"`
	NumberCreditTransactions int `json:"number_credit_transactions"`
}
