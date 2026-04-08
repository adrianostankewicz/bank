package event_test

import (
	"context"
	"testing"

	"github.com/adrianostankewicz/bank/internal/domain/account"
	"github.com/adrianostankewicz/bank/internal/domain/transaction"
	event "github.com/adrianostankewicz/bank/internal/event/account"
	usecases "github.com/adrianostankewicz/bank/internal/usecases/account"
	"github.com/stretchr/testify/assert"
)

func TestDepositEventHandlerNewAccount(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	handler := event.NewAccountDepositEventHandler(depositUC)

	result, err := handler.Handle(context.Background(), event.Event{
		Type:        "deposit",
		Destination: "100",
		Amount:      10,
	})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	response := result.(event.AccountDepositEventResponse)
	assert.Equal(t, "100", response.Destination.ID)
	assert.Equal(t, int64(10), response.Destination.Balance)
}

func TestDepositEventHandlerExistingAccount(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	handler := event.NewAccountDepositEventHandler(depositUC)

	handler.Handle(context.Background(), event.Event{
		Type:        "deposit",
		Destination: "100",
		Amount:      10,
	})

	result, err := handler.Handle(context.Background(), event.Event{
		Type:        "deposit",
		Destination: "100",
		Amount:      10,
	})

	assert.Nil(t, err)
	response := result.(event.AccountDepositEventResponse)
	assert.Equal(t, int64(20), response.Destination.Balance)
}

func TestDepositEventHandlerZeroAmount(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	handler := event.NewAccountDepositEventHandler(depositUC)

	result, err := handler.Handle(context.Background(), event.Event{
		Type:        "deposit",
		Destination: "100",
		Amount:      0,
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}
