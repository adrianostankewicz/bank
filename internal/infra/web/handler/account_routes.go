package web

import "github.com/go-chi/chi/v5"

func RegisterAccountRoutes(r chi.Router, h *WebAccountHandler) {
	r.Post("/reset", h.Reset)
	r.Get("/balance", h.Balance)
	r.Post("/event", h.Event)
}
