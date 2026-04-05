package event

import (
	"context"

	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
	usecases "github.com/adrianostankewicz/bank/internal/usecases/account"
)

type AccountOutputDTO struct {
	ID      string `json:"id"`
	Balance int64  `json:"balance"`
}

type AccountDepositEventResponse struct {
	Destination AccountOutputDTO `json:"destination"`
}

type AccountDepositEventHandler struct {
	usecase *usecases.AccountDepositUseCase
}

func NewAccountDepositEventHandler(usecase *usecases.AccountDepositUseCase) *AccountDepositEventHandler {
	return &AccountDepositEventHandler{usecase: usecase}
}

func (h *AccountDepositEventHandler) Handle(ctx context.Context, event Event) (interface{}, error) {
	output, err := h.usecase.Execute(usecases.AccountDepositInput{
		AccountID: event.Destination,
		Amount:    money.NewMoney(event.Amount),
	})
	if err != nil {
		return nil, err
	}

	return AccountDepositEventResponse{
		Destination: AccountOutputDTO{
			ID:      output.Account.ID(),
			Balance: output.Balance.Amount(),
		},
	}, nil
}
