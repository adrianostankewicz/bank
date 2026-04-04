package transaction

import (
	"time"

	"github.com/adrianostankewicz/bank/internal/core/shared/money"
	"github.com/google/uuid"
)

type TransactionType string

const (
	Deposit  TransactionType = "deposit"
	Withdraw TransactionType = "withdraw"
	Transfer TransactionType = "transfer"
)

type Transaction struct {
	id              string
	accountID       string
	amount          money.Money
	transactionType TransactionType
	createdAt       time.Time
}

func NewTransaction(accountID string, amount money.Money, transactionType TransactionType) *Transaction {
	return &Transaction{
		id:              uuid.New().String(),
		accountID:       accountID,
		amount:          amount,
		transactionType: transactionType,
		createdAt:       time.Now(),
	}
}

func (t Transaction) ID() string {
	return t.id
}

func (t Transaction) AccountID() string {
	return t.accountID
}

func (t Transaction) Amount() money.Money {
	return t.amount
}

func (t Transaction) TransactionType() TransactionType {
	return t.transactionType
}

func (t Transaction) CreatedAt() time.Time {
	return t.createdAt
}
