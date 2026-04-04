package usecases

import (
	"errors"

	"github.com/adrianostankewicz/bank/internal/core/domain/account"
	"github.com/adrianostankewicz/bank/internal/core/domain/transaction"
	"github.com/adrianostankewicz/bank/internal/core/shared/money"
)

type AccountDepositUseCase struct {
	accountRepo     account.AccountRepository
	transactionRepo transaction.TransactionRepository
}

func NewAccountDepositUseCase(
	accountRepo account.AccountRepository,
	transactionRepo transaction.TransactionRepository,
) *AccountDepositUseCase {
	return &AccountDepositUseCase{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

type AccountDepositInput struct {
	AccountID string
	Amount    money.Money
}

type AccountDepositOutput struct {
	Account     account.Account
	Transaction *transaction.Transaction
	Balance     money.Money
}

func (uc *AccountDepositUseCase) Execute(input AccountDepositInput) (*AccountDepositOutput, error) {

	if input.Amount.IsZero() {
		return nil, errors.New("amount must be greater than zero")
	}

	if input.Amount.IsNegative() {
		return nil, errors.New("amount cannot be negative")
	}

	acc, err := uc.accountRepo.FindById(input.AccountID)
	if err != nil {
		acc = account.NewBaseAccount(input.AccountID)
		err = uc.accountRepo.Save(acc)
		if err != nil {
			return nil, err
		}
	}

	tr := transaction.NewTransaction(input.AccountID, input.Amount, transaction.Deposit)

	err = uc.transactionRepo.Save(tr)
	if err != nil {
		return nil, err
	}

	transactions, err := uc.transactionRepo.FindByAccountID(input.AccountID)
	if err != nil {
		return nil, err
	}

	balance := money.NewMoney(0)
	for _, t := range transactions {
		balance = balance.Add(t.Amount())
	}

	return &AccountDepositOutput{
		Account:     acc,
		Transaction: tr,
		Balance:     balance,
	}, nil
}
