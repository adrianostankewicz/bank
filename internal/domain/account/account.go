package account

import (
	"errors"
	"time"

	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
)

type Account interface {
	ID() string
	Balance() money.Money
	Deposit(money.Money) error
	Withdraw(money.Money) error
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() *time.Time
}

type BaseAccount struct {
	id        string
	balance   money.Money
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func NewBaseAccount(id string) *BaseAccount {
	account := &BaseAccount{
		id:        id,
		balance:   money.NewMoney(0),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	return account
}

func (a BaseAccount) ID() string {
	return a.id
}

func (a BaseAccount) Balance() money.Money {
	return a.balance
}

func (a BaseAccount) CreatedAt() time.Time {
	return a.createdAt
}

func (a BaseAccount) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a BaseAccount) DeletedAt() *time.Time {
	return a.deletedAt
}

func (a *BaseAccount) Deposit(v money.Money) error {
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

func (a *BaseAccount) Withdraw(v money.Money) error {
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
