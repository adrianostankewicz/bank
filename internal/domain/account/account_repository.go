package account

import "errors"

type AccountRepository interface {
	FindById(id string) (Account, error)
	Save(account Account) error
	Reset() error
}

type InMemoryAccountRepository struct {
	accounts map[string]Account
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{
		accounts: make(map[string]Account),
	}
}

func (r *InMemoryAccountRepository) FindById(id string) (Account, error) {
	account, ok := r.accounts[id]
	if !ok {
		return nil, errors.New("account not found")
	}

	return account, nil
}

func (r *InMemoryAccountRepository) Save(account Account) error {
	r.accounts[account.ID()] = account
	return nil
}

func (r *InMemoryAccountRepository) Reset() error {
	r.accounts = make(map[string]Account)
	return nil
}
