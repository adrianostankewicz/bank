package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	event "github.com/adrianostankewicz/bank/internal/event/account"
	usecases "github.com/adrianostankewicz/bank/internal/usecases/account"
)

type WebAccountHandler struct {
	dispatcher *event.EventDispatcher
	balanceUC  *usecases.GetAccountBalanceUseCase
	resetUC    *usecases.AccountResetUseCase
}

func NewWebAccountHandler(
	dispatcher *event.EventDispatcher,
	balanceUC *usecases.GetAccountBalanceUseCase,
	resetUC *usecases.AccountResetUseCase,
) *WebAccountHandler {
	return &WebAccountHandler{
		dispatcher: dispatcher,
		balanceUC:  balanceUC,
		resetUC:    resetUC,
	}
}

func (h *WebAccountHandler) Event(w http.ResponseWriter, r *http.Request) {
	var e event.Event
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	result, err := h.dispatcher.Dispatch(r.Context(), e)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(0)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (h *WebAccountHandler) Balance(w http.ResponseWriter, r *http.Request) {
	accountID := r.URL.Query().Get("account_id")

	output, err := h.balanceUC.Execute(usecases.GetAccountBalanceInput{AccountID: accountID})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(0)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output.Balance.Amount())
}

func (h *WebAccountHandler) Reset(w http.ResponseWriter, r *http.Request) {
	if err := h.resetUC.Execute(); err != nil {
		http.Error(w, "reset failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
