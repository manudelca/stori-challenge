package domain

type MethodType string

const (
	MethodTypeDebit  MethodType = "debit"
	MethodTypeCredit MethodType = "credit"
)

type Transaction struct {
	ID         string     `json:"id"`
	Day        int        `json:"day"`
	Month      int        `json:"month"`
	MethodType MethodType `json:"method_type"`
	Amount     float64    `json:"amount"`
}
