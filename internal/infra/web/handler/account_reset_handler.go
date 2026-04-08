package web

import (
	"fmt"
	"net/http"

	usecases "github.com/adrianostankewicz/bank/internal/usecases/account"
)

type AccountResetHandler struct {
	usecase *usecases.AccountResetUseCase
}

func NewAccountResetHandler(usecase *usecases.AccountResetUseCase) *AccountResetHandler {
	return &AccountResetHandler{usecase: usecase}
}

func (h *AccountResetHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if err := h.usecase.Execute(); err != nil {
		http.Error(w, "reset failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
