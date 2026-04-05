package transaction_test

import (
	"testing"

	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
	"github.com/adrianostankewicz/bank/internal/domain/transaction"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryTransactionRepositorySave(t *testing.T) {
	repo := transaction.NewInMemoryTransactionRepository()
	tr := transaction.NewTransaction("1", money.NewMoney(100), transaction.Credit)

	err := repo.Save(tr)

	assert.Nil(t, err)
}

func TestInMemoryTransactionRepositoryFindByAccountID(t *testing.T) {
	repo := transaction.NewInMemoryTransactionRepository()
	tr := transaction.NewTransaction("1", money.NewMoney(100), transaction.Credit)
	trOne := transaction.NewTransaction("2", money.NewMoney(25), transaction.Credit)
	trTwo := transaction.NewTransaction("1", money.NewMoney(120), transaction.Debit)

	repo.Save(tr)
	repo.Save(trOne)
	repo.Save(trTwo)

	result, err := repo.FindByAccountID("1")

	assert.Nil(t, err)
	assert.Len(t, result, 2)
}

func TestInMemoryTransactionRepositoryReset(t *testing.T) {
	repo := transaction.NewInMemoryTransactionRepository()
	tr := transaction.NewTransaction("1", money.NewMoney(100), transaction.Credit)

	repo.Save(tr)

	repo.Reset()
	_, err := repo.FindByAccountID(tr.AccountID())

	assert.Error(t, err)
}
