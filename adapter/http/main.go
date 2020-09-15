package http

import (
	"net/http"
	"github.com/g-barros/expert-session-finance/adapter/http/transaction"
	"github.com/g-barros/expert-session-finance/adapter/http/actuator"
)
// Init inicia servidor e rotas
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}