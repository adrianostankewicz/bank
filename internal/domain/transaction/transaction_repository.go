package transaction

import "errors"

type TransactionRepository interface {
	Save(transaction *Transaction) error
	FindByAccountID(accountID string) ([]*Transaction, error)
	Reset() error
}

type InMemoryTransactionRepository struct {
	transactions map[string]*Transaction
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{
		transactions: make(map[string]*Transaction),
	}
}

func (r *InMemoryTransactionRepository) Save(t *Transaction) error {
	r.transactions[t.ID()] = t
	return nil
}

func (r *InMemoryTransactionRepository) FindByAccountID(accountID string) ([]*Transaction, error) {
	var result []*Transaction
	for _, t := range r.transactions {
		if t.AccountID() == accountID {
			result = append(result, t)
		}
	}

	if len(result) == 0 {
		return nil, errors.New("no records found")
	}
	return result, nil
}

func (r *InMemoryTransactionRepository) Reset() error {
	r.transactions = make(map[string]*Transaction)
	return nil
}
