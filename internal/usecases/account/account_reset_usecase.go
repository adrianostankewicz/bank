package usecases

import (
	"github.com/adrianostankewicz/bank/internal/domain/account"
	"github.com/adrianostankewicz/bank/internal/domain/transaction"
)

type AccountResetUseCase struct {
	accountRepo     account.AccountRepository
	transactionRepo transaction.TransactionRepository
}

func NewAccountResetUseCase(
	accountRepo account.AccountRepository,
	transactionRepo transaction.TransactionRepository,
) *AccountResetUseCase {
	return &AccountResetUseCase{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

func (uc *AccountResetUseCase) Execute() error {
	err := uc.accountRepo.Reset()
	if err != nil {
		return err
	}

	err = uc.transactionRepo.Reset()
	if err != nil {
		return err
	}

	return nil
}
