package usecases

import (
	"errors"

	"github.com/adrianostankewicz/bank/internal/core/domain/account"
	"github.com/adrianostankewicz/bank/internal/core/domain/transaction"
	"github.com/adrianostankewicz/bank/internal/core/shared/money"
)

type AccountWithdrawUseCase struct {
	accountRepo     account.AccountRepository
	transactionRepo transaction.TransactionRepository
}

func NewAccountWithdrawUseCase(
	accountRepo account.AccountRepository,
	transactionRepo transaction.TransactionRepository,
) *AccountWithdrawUseCase {
	return &AccountWithdrawUseCase{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

type AccountWithdrawInput struct {
	AccountID string
	Amount    money.Money
}

type AccountWithdrawOutput struct {
	Account     account.Account
	Transaction *transaction.Transaction
	Balance     money.Money
}

func (uc *AccountWithdrawUseCase) Execute(input AccountWithdrawInput) (*AccountWithdrawOutput, error) {
	if input.Amount.IsZero() {
		return nil, errors.New("amount must be greater than zero")
	}

	if input.Amount.IsNegative() {
		return nil, errors.New("amount cannot be negative")
	}

	acc, err := uc.accountRepo.FindById(input.AccountID)
	if err != nil {
		return nil, errors.New("account not found")
	}

	transactions, err := uc.transactionRepo.FindByAccountID(input.AccountID)
	if err != nil {
		return nil, err
	}

	balance := transaction.CalculateBalance(transactions)

	if input.Amount.IsGreaterThan(balance) {
		return nil, errors.New("insufficient balance")
	}

	tr := transaction.NewTransaction(input.AccountID, input.Amount.Negative(), transaction.Debit)

	if err := uc.transactionRepo.Save(tr); err != nil {
		return nil, err
	}

	transactions, err = uc.transactionRepo.FindByAccountID(input.AccountID)
	if err != nil {
		return nil, err
	}

	balance = transaction.CalculateBalance(transactions)

	return &AccountWithdrawOutput{
		Account:     acc,
		Transaction: tr,
		Balance:     balance,
	}, nil
}
