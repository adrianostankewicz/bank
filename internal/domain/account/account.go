package account

import (
	"errors"
	"time"

	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
	"github.com/google/uuid"
)

type Account struct {
	id        string
	balance   money.Money
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func NewAccount() *Account {
	account := &Account{
		id:        uuid.New().String(),
		balance:   money.NewMoney(0),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	return account
}

func (a Account) ID() string {
	return a.id
}

func (a Account) Balance() money.Money {
	return a.balance
}

func (a Account) CreatedAt() time.Time {
	return a.createdAt
}

func (a Account) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a Account) DeletedAt() *time.Time {
	return a.deletedAt
}

func (a *Account) Deposit(v money.Money) error {
	if v.IsZero() {
		return errors.New("amount must be greater than zero")
	}

	if v.IsNegative() {
		return errors.New("amount cannot be negative")
	}

	a.balance = a.Balance().Add(v)
	a.updatedAt = time.Now()

	return nil
}

func (a *Account) Withdraw(v money.Money) error {
	if v.IsNegative() {
		return errors.New("amount cannot be negative")
	}

	if v.IsZero() {
		return errors.New("amount must be greater than zero")
	}

	newBalance := a.balance.Sub(v)
	if newBalance.IsNegative() {
		return errors.New("insufficient balance")
	}

	a.balance = a.Balance().Sub(v)
	a.updatedAt = time.Now()

	return nil
}
