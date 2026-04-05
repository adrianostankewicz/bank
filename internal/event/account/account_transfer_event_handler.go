package event

import (
	"context"

	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
	usecases "github.com/adrianostankewicz/bank/internal/usecases/account"
)

type AccountTransferOutputDTO struct {
	ID      string `json:"id"`
	Balance int64  `json:"balance"`
}

type AccountTransferEventResponse struct {
	Origin      AccountTransferOutputDTO `json:"origin"`
	Destination AccountTransferOutputDTO `json:"destination"`
}

type AccountTransferEventHandler struct {
	usecase *usecases.AccountTransferUseCase
}

func NewAccountTransferEventHandler(usecase *usecases.AccountTransferUseCase) *AccountTransferEventHandler {
	return &AccountTransferEventHandler{usecase: usecase}
}

func (h *AccountTransferEventHandler) Handle(ctx context.Context, event Event) (interface{}, error) {
	output, err := h.usecase.Execute(usecases.AccountTransferInput{
		OriginID:      event.Origin,
		DestinationID: event.Destination,
		Amount:        money.NewMoney(event.Amount),
	})
	if err != nil {
		return nil, err
	}

	return AccountTransferEventResponse{
		Origin: AccountTransferOutputDTO{
			ID:      output.Origin.ID(),
			Balance: output.OriginBalance.Amount(),
		},
		Destination: AccountTransferOutputDTO{
			ID:      output.Destination.ID(),
			Balance: output.DestinationBalance.Amount(),
		},
	}, nil
}
