package repository

import "github.com/manudelca/stori-challenge/internal/domain"

type TransactionRepository interface {
	SaveTransaction(transaction domain.Transaction)
}

type transactionRepository struct {
	// Esto viene a representar un storage. Preferentemente de tipo NoSQL para mayor escalabilidad
	storage map[string]domain.Transaction
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{
		storage: make(map[string]domain.Transaction),
	}
}

func (r *transactionRepository) SaveTransaction(transaction domain.Transaction) {
	r.storage[transaction.ID] = transaction
}
