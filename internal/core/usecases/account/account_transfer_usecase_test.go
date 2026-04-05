package usecases_test

import (
	"testing"

	"github.com/adrianostankewicz/bank/internal/core/domain/account"
	"github.com/adrianostankewicz/bank/internal/core/domain/transaction"
	"github.com/adrianostankewicz/bank/internal/core/shared/money"
	usecases "github.com/adrianostankewicz/bank/internal/core/usecases/account"
	"github.com/stretchr/testify/assert"
)

func TestTransferUseCaseSuccess(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	transferUC := usecases.NewAccountTransferUseCase(accountRepo, transactionRepo)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(200)})
	depositUC.Execute(usecases.AccountDepositInput{AccountID: "300", Amount: money.NewMoney(100)})

	output, err := transferUC.Execute(usecases.AccountTransferInput{
		OriginID:      "100",
		DestinationID: "300",
		Amount:        money.NewMoney(50),
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, int64(150), output.OriginBalance.Amount())
	assert.Equal(t, int64(150), output.DestinationBalance.Amount())
}

func TestTransferUseCaseOriginNotFound(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	transferUC := usecases.NewAccountTransferUseCase(accountRepo, transactionRepo)

	output, err := transferUC.Execute(usecases.AccountTransferInput{
		OriginID:      "200",
		DestinationID: "300",
		Amount:        money.NewMoney(50),
	})

	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestTransferUseCaseDestinationNotFound(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	transferUC := usecases.NewAccountTransferUseCase(accountRepo, transactionRepo)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(200)})

	output, err := transferUC.Execute(usecases.AccountTransferInput{
		OriginID:      "100",
		DestinationID: "300",
		Amount:        money.NewMoney(50),
	})

	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestTransferUseCaseInsufficientBalance(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	transferUC := usecases.NewAccountTransferUseCase(accountRepo, transactionRepo)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(50)})
	depositUC.Execute(usecases.AccountDepositInput{AccountID: "300", Amount: money.NewMoney(50)})

	output, err := transferUC.Execute(usecases.AccountTransferInput{
		OriginID:      "100",
		DestinationID: "300",
		Amount:        money.NewMoney(100),
	})

	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestTransferUseCaseZeroAmount(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	transferUC := usecases.NewAccountTransferUseCase(accountRepo, transactionRepo)

	output, err := transferUC.Execute(usecases.AccountTransferInput{
		OriginID:      "100",
		DestinationID: "300",
		Amount:        money.NewMoney(0),
	})

	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestTransferUseCaseNegativeAmount(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	transferUC := usecases.NewAccountTransferUseCase(accountRepo, transactionRepo)

	output, err := transferUC.Execute(usecases.AccountTransferInput{
		OriginID:      "100",
		DestinationID: "300",
		Amount:        money.NewMoney(-50),
	})

	assert.Error(t, err)
	assert.Nil(t, output)
}
