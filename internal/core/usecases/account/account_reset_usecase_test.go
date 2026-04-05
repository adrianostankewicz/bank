package usecases_test

import (
	"testing"

	"github.com/adrianostankewicz/bank/internal/core/domain/account"
	"github.com/adrianostankewicz/bank/internal/core/domain/transaction"
	"github.com/adrianostankewicz/bank/internal/core/shared/money"
	usecases "github.com/adrianostankewicz/bank/internal/core/usecases/account"
	"github.com/stretchr/testify/assert"
)

func TestResetUseCaseSuccess(t *testing.T) {
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	resetUC := usecases.NewAccountResetUseCase(accountRepo, transactionRepo)

	depositUC.Execute(usecases.AccountDepositInput{AccountID: "100", Amount: money.NewMoney(200)})

	err := resetUC.Execute()

	assert.Nil(t, err)

	_, err = accountRepo.FindById("100")
	assert.Error(t, err)

	transactions, _ := transactionRepo.FindByAccountID("100")
	assert.Len(t, transactions, 0)
}
