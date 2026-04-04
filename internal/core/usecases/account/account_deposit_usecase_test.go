package usecases_test

import (
	"testing"

	"github.com/adrianostankewicz/bank/internal/core/domain/account"
	"github.com/adrianostankewicz/bank/internal/core/domain/transaction"
	"github.com/adrianostankewicz/bank/internal/core/shared/money"
	usecases "github.com/adrianostankewicz/bank/internal/core/usecases/account"
	"github.com/stretchr/testify/assert"
)

func TestAccountDepositUseCaseNewAccount(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	uc := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)

	input := usecases.AccountDepositInput{
		AccountID: "100",
		Amount:    money.NewMoney(100),
	}

	output, err := uc.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "100", output.Account.ID())
	assert.Equal(t, int64(100), output.Balance.Amount())
	assert.NotNil(t, output.Transaction)
}

func TestAccountDepositUseCaseExistingAccount(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	uc := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)

	input := usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(100)}
	uc.Execute(input)

	inputTwo := usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(50)}
	output, err := uc.Execute(inputTwo)

	assert.Nil(t, err)
	assert.Equal(t, int64(150), output.Balance.Amount())
}

func TestDepositUseCaseZeroAmount(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	uc := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)

	input := usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(0)}
	output, err := uc.Execute(input)

	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestDepositUseCaseNegativeAmount(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	uc := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)

	input := usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(-100)}
	output, err := uc.Execute(input)

	assert.Error(t, err)
	assert.Nil(t, output)
}
