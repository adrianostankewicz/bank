package usecases

import (
	"errors"

	"github.com/adrianostankewicz/bank/internal/domain/account"
	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
	"github.com/adrianostankewicz/bank/internal/domain/transaction"
)

type GetAccountBalanceUseCase struct {
	accountRepo     account.AccountRepository
	transactionRepo transaction.TransactionRepository
}

func NewGetAccountBalanceUseCase(
	accountRepo account.AccountRepository,
	transactionRepo transaction.TransactionRepository,
) *GetAccountBalanceUseCase {
	return &GetAccountBalanceUseCase{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

type GetAccountBalanceInput struct {
	AccountID string
}

type GetAccountBalanceOutput struct {
	Account account.Account
	Balance money.Money
}

func (uc *GetAccountBalanceUseCase) Execute(input GetAccountBalanceInput) (*GetAccountBalanceOutput, error) {
	acc, err := uc.accountRepo.FindById(input.AccountID)
	if err != nil {
		return nil, errors.New("account not found")
	}

	transactions, err := uc.transactionRepo.FindByAccountID(input.AccountID)
	if err != nil {
		return nil, err
	}

	balance := transaction.CalculateBalance(transactions)

	return &GetAccountBalanceOutput{
		Account: acc,
		Balance: balance,
	}, nil
}
