package transaction_test

import (
	"testing"

	"github.com/adrianostankewicz/bank/internal/core/domain/transaction"
	"github.com/adrianostankewicz/bank/internal/core/shared/money"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryTransactionRepositorySave(t *testing.T) {
	repo := transaction.NewInMemoryTransactionRepository()
	tr := transaction.NewTransaction("1", money.NewMoney(100), transaction.Deposit)

	err := repo.Save(tr)

	assert.Nil(t, err)
}

func TestInMemoryTransactionRepositoryFindByAccountID(t *testing.T) {
	repo := transaction.NewInMemoryTransactionRepository()
	tr := transaction.NewTransaction("1", money.NewMoney(100), transaction.Deposit)
	trOne := transaction.NewTransaction("2", money.NewMoney(25), transaction.Deposit)
	trTwo := transaction.NewTransaction("1", money.NewMoney(120), transaction.Withdraw)

	repo.Save(tr)
	repo.Save(trOne)
	repo.Save(trTwo)

	result, err := repo.FindByAccountID("1")

	assert.Nil(t, err)
	assert.Len(t, result, 2)
}

func TestInMemoryTransactionRepositoryReset(t *testing.T) {
	repo := transaction.NewInMemoryTransactionRepository()
	tr := transaction.NewTransaction("1", money.NewMoney(100), transaction.Deposit)

	repo.Save(tr)

	repo.Reset()
	_, err := repo.FindByAccountID(tr.AccountID())

	assert.Error(t, err)
}
