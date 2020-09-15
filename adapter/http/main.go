package http

import (
	"github.com/g-barros/expert-session-finance/adapter/http/actuator"
	"github.com/g-barros/expert-session-finance/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// Init inicia servidor e rotas
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
