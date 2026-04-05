package transaction_test

import (
	"testing"

	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
	"github.com/adrianostankewicz/bank/internal/domain/transaction"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	tr := transaction.NewTransaction("1", money.NewMoney(100), transaction.Credit)

	assert.NotEmpty(t, tr.ID())
	assert.Equal(t, "1", tr.AccountID())
	assert.Equal(t, int64(100), tr.Amount().Amount())
	assert.Equal(t, transaction.Credit, tr.TransactionType())
	assert.NotZero(t, tr.CreatedAt())
}
