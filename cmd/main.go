// cmd/main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/adrianostankewicz/bank/internal/domain/account"
	"github.com/adrianostankewicz/bank/internal/domain/transaction"
	event "github.com/adrianostankewicz/bank/internal/event/account"
	web "github.com/adrianostankewicz/bank/internal/infra/web/handler"
	usecases "github.com/adrianostankewicz/bank/internal/usecases/account"
	"github.com/go-chi/chi/v5"
)

func main() {
	// repositories
	accountRepo := account.NewInMemoryAccountRepository()
	transactionRepo := transaction.NewInMemoryTransactionRepository()

	// usecases
	depositUC := usecases.NewAccountDepositUseCase(accountRepo, transactionRepo)
	withdrawUC := usecases.NewAccountWithdrawUseCase(accountRepo, transactionRepo)
	transferUC := usecases.NewAccountTransferUseCase(accountRepo, transactionRepo)
	balanceUC := usecases.NewGetAccountBalanceUseCase(accountRepo, transactionRepo)
	resetUC := usecases.NewAccountResetUseCase(accountRepo, transactionRepo)

	// event handlers
	depositEventHandler := event.NewAccountDepositEventHandler(depositUC)
	withdrawEventHandler := event.NewAccountWithdrawEventHandler(withdrawUC)
	transferEventHandler := event.NewAccountTransferEventHandler(transferUC)

	// dispatcher
	dispatcher := event.NewEventDispatcher()
	dispatcher.Register("deposit", depositEventHandler)
	dispatcher.Register("withdraw", withdrawEventHandler)
	dispatcher.Register("transfer", transferEventHandler)

	// event handler
	eventHandler := web.NewAccountEventHandler(dispatcher)

	// http handler
	balanceHandler := web.NewAccountBalanceHandler(balanceUC)
	resetHandler := web.NewAccountResetHandler(resetUC)

	// router
	r := chi.NewRouter()
	web.RegisterAccountRoutes(r, eventHandler, balanceHandler, resetHandler)

	// server
	fmt.Println("server running on port 3000")
	http.ListenAndServe(":3000", r)
}
