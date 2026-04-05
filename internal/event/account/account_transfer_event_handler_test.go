package event_test

import (
	"context"
	"testing"

	"github.com/adrianostankewicz/bank/internal/domain/account"
	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
	"github.com/adrianostankewicz/bank/internal/domain/transaction"
	event "github.com/adrianostankewicz/bank/internal/event/account"
	usecases "github.com/adrianostankewicz/bank/internal/usecases/account"
	"github.com/stretchr/testify/assert"
)

func TestTransferEventHandlerSuccess(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	transferUC := usecases.NewAccountTransferUseCase(accountRepo, transactionRepo)
	handler := event.NewAccountTransferEventHandler(transferUC)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(200)})
	accountRepo.Save(account.NewBaseAccount("300"))

	result, err := handler.Handle(context.Background(), event.Event{
		Type:        "transfer",
		Origin:      "100",
		Destination: "300",
		Amount:      50,
	})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	response := result.(event.AccountTransferEventResponse)
	assert.Equal(t, "100", response.Origin.ID)
	assert.Equal(t, int64(150), response.Origin.Balance)
	assert.Equal(t, "300", response.Destination.ID)
	assert.Equal(t, int64(50), response.Destination.Balance)
}

func TestTransferEventHandlerOriginNotFound(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	transferUC := usecases.NewAccountTransferUseCase(accountRepo, transactionRepo)
	handler := event.NewAccountTransferEventHandler(transferUC)

	result, err := handler.Handle(context.Background(), event.Event{
		Type:        "transfer",
		Origin:      "200",
		Destination: "300",
		Amount:      50,
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestTransferEventHandlerDestinationNotFound(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	transferUC := usecases.NewAccountTransferUseCase(accountRepo, transactionRepo)
	handler := event.NewAccountTransferEventHandler(transferUC)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(200)})

	result, err := handler.Handle(context.Background(), event.Event{
		Type:        "transfer",
		Origin:      "100",
		Destination: "300",
		Amount:      50,
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestTransferEventHandlerInsufficientBalance(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	transferUC := usecases.NewAccountTransferUseCase(accountRepo, transactionRepo)
	handler := event.NewAccountTransferEventHandler(transferUC)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(50)})
	accountRepo.Save(account.NewBaseAccount("300"))

	result, err := handler.Handle(context.Background(), event.Event{
		Type:        "transfer",
		Origin:      "100",
		Destination: "300",
		Amount:      100,
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}
