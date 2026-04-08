package usecases_test

import (
	"testing"

	"github.com/adrianostankewicz/bank/internal/domain/account"
	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
	"github.com/adrianostankewicz/bank/internal/domain/transaction"
	usecases "github.com/adrianostankewicz/bank/internal/usecases/account"
	"github.com/stretchr/testify/assert"
)

func TestAccountWithdrawUseCaseSuccess(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	withdrawUC := usecases.NewAccountWithdrawUseCase(accountRepo, transactionRepo)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(200)})

	output, err := withdrawUC.Execute(usecases.AccountWithdrawInput{AccountID: "100", Amount: money.NewMoney(50)})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, int64(150), output.Balance.Amount())
}

func TestAccountWithdrawUseCaseAccountNotFound(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	withdrawUC := usecases.NewAccountWithdrawUseCase(accountRepo, transactionRepo)

	output, err := withdrawUC.Execute(usecases.AccountWithdrawInput{AccountID: "200", Amount: money.NewMoney(50)})

	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestAccountWithdrawUseCaseInsufficientBalance(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	withdrawUC := usecases.NewAccountWithdrawUseCase(accountRepo, transactionRepo)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(50)})

	output, err := withdrawUC.Execute(usecases.AccountWithdrawInput{AccountID: "100", Amount: money.NewMoney(100)})

	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestAccountWithdrawUseCaseZeroAmount(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	withdrawUC := usecases.NewAccountWithdrawUseCase(accountRepo, transactionRepo)

	output, err := withdrawUC.Execute(usecases.AccountWithdrawInput{AccountID: "100", Amount: money.NewMoney(0)})

	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestAccountWithdrawUseCaseNegativeAmount(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	withdrawUC := usecases.NewAccountWithdrawUseCase(accountRepo, transactionRepo)

	output, err := withdrawUC.Execute(usecases.AccountWithdrawInput{AccountID: "100", Amount: money.NewMoney(-50)})

	assert.Error(t, err)
	assert.Nil(t, output)
}
