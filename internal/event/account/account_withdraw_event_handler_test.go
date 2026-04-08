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

func TestWithdrawEventHandlerSuccess(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	withdrawUC := usecases.NewAccountWithdrawUseCase(accountRepo, transactionRepo)
	handler := event.NewAccountWithdrawEventHandler(withdrawUC)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(200)})

	result, err := handler.Handle(context.Background(), event.Event{
		Type:   "withdraw",
		Origin: "100",
		Amount: 50,
	})

	assert.Nil(t, err)
	assert.NotNil(t, result)

	response := result.(event.AccountWithdrawEventResponse)
	assert.Equal(t, "100", response.Origin.ID)
	assert.Equal(t, int64(150), response.Origin.Balance)
}

func TestWithdrawEventHandlerAccountNotFound(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	withdrawUC := usecases.NewAccountWithdrawUseCase(accountRepo, transactionRepo)
	handler := event.NewAccountWithdrawEventHandler(withdrawUC)

	result, err := handler.Handle(context.Background(), event.Event{
		Type:   "withdraw",
		Origin: "200",
		Amount: 50,
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestWithdrawEventHandlerInsufficientBalance(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	withdrawUC := usecases.NewAccountWithdrawUseCase(accountRepo, transactionRepo)
	handler := event.NewAccountWithdrawEventHandler(withdrawUC)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(50)})

	result, err := handler.Handle(context.Background(), event.Event{
		Type:   "withdraw",
		Origin: "100",
		Amount: 100,
	})

	assert.Error(t, err)
	assert.Nil(t, result)
}
