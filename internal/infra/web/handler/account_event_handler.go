package web

import (
	"encoding/json"
	"net/http"

	event "github.com/adrianostankewicz/bank/internal/event/account"
)

type AccountEventHandler struct {
	dispatcher *event.EventDispatcher
}

func NewAccountEventHandler(dispatcher *event.EventDispatcher) *AccountEventHandler {
	return &AccountEventHandler{dispatcher: dispatcher}
}

func (h *AccountEventHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var e event.Event
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
