package usecases

import (
	"errors"

	"github.com/adrianostankewicz/bank/internal/core/domain/account"
	"github.com/adrianostankewicz/bank/internal/core/domain/transaction"
	"github.com/adrianostankewicz/bank/internal/core/shared/money"
)

type AccountTransferUseCase struct {
	accountRepo     account.AccountRepository
	transactionRepo transaction.TransactionRepository
}

func NewAccountTransferUseCase(
	accountRepo account.AccountRepository,
	transactionRepo transaction.TransactionRepository,
) *AccountTransferUseCase {
	return &AccountTransferUseCase{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

type AccountTransferInput struct {
	OriginID      string
	DestinationID string
	Amount        money.Money
}

type AccountTransferOutput struct {
	Origin             account.Account
	Destination        account.Account
	OriginBalance      money.Money
	DestinationBalance money.Money
}

func (uc *AccountTransferUseCase) Execute(input AccountTransferInput) (*AccountTransferOutput, error) {
	if input.Amount.IsZero() {
		return nil, errors.New("amount must be greater than zero")
	}

	if input.Amount.IsNegative() {
		return nil, errors.New("amount cannot be negative")
	}

	origin, err := uc.accountRepo.FindById(input.OriginID)
	if err != nil {
		return nil, errors.New("origin account not found")
	}

	destination, err := uc.accountRepo.FindById(input.DestinationID)
	if err != nil {
		return nil, errors.New("destination account not found")
	}

	originTransactions, err := uc.transactionRepo.FindByAccountID(input.OriginID)
	if err != nil {
		return nil, err
	}

	originBalance := money.NewMoney(0)
	for _, t := range originTransactions {
		originBalance = originBalance.Add(t.Amount())
	}

	if input.Amount.IsGreaterThan(originBalance) {
		return nil, errors.New("insufficient balance")
	}

	debit := transaction.NewTransaction(input.OriginID, input.Amount, transaction.Debit)
	err = uc.transactionRepo.Save(debit)
	if err != nil {
		return nil, err
	}

	credit := transaction.NewTransaction(input.DestinationID, input.Amount, transaction.Credit)
	err = uc.transactionRepo.Save(credit)
	if err != nil {
		compensation := transaction.NewTransaction(input.OriginID, input.Amount, transaction.Credit)
		uc.transactionRepo.Save(compensation)
		return nil, err
	}

	destinationTransactions, err := uc.transactionRepo.FindByAccountID(input.DestinationID)
	if err != nil {
		return nil, err
	}

	originBalance = originBalance.Sub(input.Amount)

	destinationBalance := money.NewMoney(0)
	for _, t := range destinationTransactions {
		destinationBalance = destinationBalance.Add(t.Amount())
	}

	return &AccountTransferOutput{
		Origin:             origin,
		Destination:        destination,
		OriginBalance:      originBalance,
		DestinationBalance: destinationBalance,
	}, nil
}
