package event

import (
	"context"

	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
	usecases "github.com/adrianostankewicz/bank/internal/usecases/account"
)

type AccountWithdrawOutputDTO struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type AccountWithdrawEventResponse struct {
	Origin AccountWithdrawOutputDTO `json:"origin"`
}

type AccountWithdrawEventHandler struct {
	usecase *usecases.AccountWithdrawUseCase
}

func NewAccountWithdrawEventHandler(usecase *usecases.AccountWithdrawUseCase) *AccountWithdrawEventHandler {
	return &AccountWithdrawEventHandler{usecase: usecase}
}

func (h *AccountWithdrawEventHandler) Handle(ctx context.Context, event Event) (interface{}, error) {
	output, err := h.usecase.Execute(usecases.AccountWithdrawInput{
		AccountID: event.Origin,
		Amount:    money.NewMoneyFromFloat(event.Amount),
	})
	if err != nil {
		return nil, err
	}

	return AccountWithdrawEventResponse{
		Origin: AccountWithdrawOutputDTO{
			ID:      output.Account.ID(),
			Balance: output.Balance.ToFloat(),
		},
	}, nil
}
