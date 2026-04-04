package account_test

import (
	"testing"
	"time"

	"github.com/adrianostankewicz/bank/internal/domain/account"
	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	a := account.NewAccount()

	assert.NotEmpty(t, a.ID())
	assert.Equal(t, int64(0), a.Balance().Amount())
	assert.NotZero(t, a.CreatedAt())
	assert.NotZero(t, a.UpdatedAt())
	assert.Nil(t, a.DeletedAt())
}

func TestAccountDeposit(t *testing.T) {
	a := account.NewAccount()
	created := a.CreatedAt()
	updated := a.UpdatedAt()
	id := a.ID()
	balance := a.Balance()

	time.Sleep(time.Millisecond)

	v := money.NewMoney(100)
	a.Deposit(v)

	assert.Equal(t, a.ID(), id)
	assert.Greater(t, a.Balance().Amount(), balance.Amount())
	assert.Equal(t, int64(100), a.Balance().Amount())
	assert.True(t, a.UpdatedAt().After(updated))
	assert.True(t, a.CreatedAt().Equal(created))
	assert.Nil(t, a.DeletedAt())
}

func TestAccountDepositZeroValue(t *testing.T) {
	a := account.NewAccount()

	v := money.NewMoney(0)
	err := a.Deposit(v)

	assert.Error(t, err)
}

func TestAccountDepositNegativeValue(t *testing.T) {
	a := account.NewAccount()

	v := money.NewMoney(-100)
	err := a.Deposit(v)

	assert.Error(t, err)
}

func TestAccountWithdraw(t *testing.T) {
	a := account.NewAccount()
	v := money.NewMoney(200)
	a.Deposit(v)

	id := a.ID()
	balance := a.Balance()
	created := a.CreatedAt()
	updated := a.UpdatedAt()

	time.Sleep(time.Millisecond)

	v = money.NewMoney(100)
	a.Withdraw(v)

	assert.Equal(t, a.ID(), id)
	assert.Less(t, a.Balance().Amount(), balance.Amount())
	assert.Equal(t, int64(100), a.Balance().Amount())
	assert.True(t, a.UpdatedAt().After(updated))
	assert.True(t, a.CreatedAt().Equal(created))
	assert.Nil(t, a.DeletedAt())
}

func TestAccountWithdrawNegativeValue(t *testing.T) {
	a := account.NewAccount()
	v := money.NewMoney(50)
	a.Deposit(v)

	time.Sleep(time.Millisecond)

	v = money.NewMoney(-100)
	err := a.Withdraw(v)

	assert.Error(t, err)
}

func TestAccountWithdrawZeroValue(t *testing.T) {
	a := account.NewAccount()
	v := money.NewMoney(50)
	a.Deposit(v)

	time.Sleep(time.Millisecond)

	v = money.NewMoney(0)
	err := a.Withdraw(v)

	assert.Error(t, err)
}

func TestAccountWithdrawInsufficientBalance(t *testing.T) {
	a := account.NewAccount()
	v := money.NewMoney(50)
	a.Deposit(v)

	time.Sleep(time.Millisecond)

	v = money.NewMoney(100)
	err := a.Withdraw(v)

	assert.Error(t, err)
}
