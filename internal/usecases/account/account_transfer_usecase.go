package usecases

import (
	"errors"

	"github.com/adrianostankewicz/bank/internal/domain/account"
	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
	"github.com/adrianostankewicz/bank/internal/domain/transaction"
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
		destination = account.NewBaseAccount(input.DestinationID)
		if err := uc.accountRepo.Save(destination); err != nil {
			return nil, err
		}
	}

	originTransactions, err := uc.transactionRepo.FindByAccountID(input.OriginID)
	if err != nil {
		return nil, err
	}

	originBalance := transaction.CalculateBalance(originTransactions)

	if input.Amount.IsGreaterThan(originBalance) {
		return nil, errors.New("insufficient balance")
	}

	debit := transaction.NewTransaction(input.OriginID, input.Amount.Negative(), transaction.Debit)
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

	originTransactions, err = uc.transactionRepo.FindByAccountID(input.OriginID)
	if err != nil {
		return nil, err
	}

	originBalance = transaction.CalculateBalance(originTransactions)

	destinationTransactions, err := uc.transactionRepo.FindByAccountID(input.DestinationID)
	if err != nil {
		return nil, err
	}

	destinationBalance := transaction.CalculateBalance(destinationTransactions)

	return &AccountTransferOutput{
		Origin:             origin,
		Destination:        destination,
		OriginBalance:      originBalance,
		DestinationBalance: destinationBalance,
	}, nil
}
