package transaction_test

import (
	"testing"

	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
	"github.com/adrianostankewicz/bank/internal/domain/transaction"
	"github.com/stretchr/testify/assert"
)

func TestCalculateBalanceWithDeposits(t *testing.T) {
	transactions := []*transaction.Transaction{
		transaction.NewTransaction("100", money.NewMoney(100), transaction.Credit),
		transaction.NewTransaction("100", money.NewMoney(50), transaction.Credit),
	}

	balance := transaction.CalculateBalance(transactions)

	assert.Equal(t, int64(150), balance.Amount())
}

func TestCalculateBalanceWithWithdraw(t *testing.T) {
	transactions := []*transaction.Transaction{
		transaction.NewTransaction("100", money.NewMoney(100), transaction.Credit),
		transaction.NewTransaction("100", money.NewMoney(-50), transaction.Debit),
	}

	balance := transaction.CalculateBalance(transactions)

	assert.Equal(t, int64(50), balance.Amount())
}

func TestCalculateBalanceEmpty(t *testing.T) {
	balance := transaction.CalculateBalance([]*transaction.Transaction{})

	assert.Equal(t, int64(0), balance.Amount())
}

func TestCalculateBalanceZero(t *testing.T) {
	transactions := []*transaction.Transaction{
		transaction.NewTransaction("100", money.NewMoney(100), transaction.Credit),
		transaction.NewTransaction("100", money.NewMoney(-100), transaction.Debit),
	}

	balance := transaction.CalculateBalance(transactions)

	assert.Equal(t, int64(0), balance.Amount())
}
