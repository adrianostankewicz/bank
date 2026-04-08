package web

import "github.com/go-chi/chi/v5"

func RegisterAccountRoutes(r chi.Router,
	eventHandler *AccountEventHandler,
	balanceHandler *AccountBalanceHandler,
	resetHandler *AccountResetHandler,
) {
	r.Post("/reset", resetHandler.Handle)
	r.Get("/balance", balanceHandler.Handle)
	r.Post("/event", eventHandler.Handle)
}
