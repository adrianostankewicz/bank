package web

import (
	"encoding/json"
	"net/http"

	usecases "github.com/adrianostankewicz/bank/internal/usecases/account"
)

type AccountBalanceHandler struct {
	usecase *usecases.GetAccountBalanceUseCase
}

func NewAccountBalanceHandler(usecase *usecases.GetAccountBalanceUseCase) *AccountBalanceHandler {
	return &AccountBalanceHandler{usecase: usecase}
}

func (h *AccountBalanceHandler) Handle(w http.ResponseWriter, r *http.Request) {
	accountID := r.URL.Query().Get("account_id")

	output, err := h.usecase.Execute(usecases.GetAccountBalanceInput{AccountID: accountID})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(0)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output.Balance.ToFloat())
}
