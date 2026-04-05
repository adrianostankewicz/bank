package usecases_test

import (
	"testing"

	"github.com/adrianostankewicz/bank/internal/core/domain/account"
	"github.com/adrianostankewicz/bank/internal/core/domain/transaction"
	"github.com/adrianostankewicz/bank/internal/core/shared/money"
	usecases "github.com/adrianostankewicz/bank/internal/core/usecases/account"
	"github.com/stretchr/testify/assert"
)

func TestGetBalanceUseCaseSuccess(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	getBalanceUC := usecases.NewGetAccountBalanceUseCase(accountRepo, transactionRepo)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(200)})

	output, err := getBalanceUC.Execute(usecases.GetAccountBalanceInput{AccountID: "100"})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, int64(200), output.Balance.Amount())
}

func TestGetBalanceUseCaseAccountNotFound(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	getBalanceUC := usecases.NewGetAccountBalanceUseCase(accountRepo, transactionRepo)

	output, err := getBalanceUC.Execute(usecases.GetAccountBalanceInput{AccountID: "100"})

	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestGetBalanceUseCaseAfterWithdraw(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	withdrawUC := usecases.NewAccountWithdrawUseCase(accountRepo, transactionRepo)
	getBalanceUC := usecases.NewGetAccountBalanceUseCase(accountRepo, transactionRepo)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(200)})
	withdrawUC.Execute(usecases.AccountWithdrawInput{AccountID: "100", Amount: money.NewMoney(50)})

	output, err := getBalanceUC.Execute(usecases.GetAccountBalanceInput{AccountID: "100"})

	assert.Nil(t, err)
	assert.Equal(t, int64(150), output.Balance.Amount())
}

func TestGetBalanceUseCaseZeroBalance(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	withdrawUC := usecases.NewAccountWithdrawUseCase(accountRepo, transactionRepo)
	getBalanceUC := usecases.NewGetAccountBalanceUseCase(accountRepo, transactionRepo)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(200)})
	withdrawUC.Execute(usecases.AccountWithdrawInput{AccountID: "100", Amount: money.NewMoney(200)})

	output, err := getBalanceUC.Execute(usecases.GetAccountBalanceInput{AccountID: "100"})

	assert.Nil(t, err)
	assert.Equal(t, int64(0), output.Balance.Amount())
}
