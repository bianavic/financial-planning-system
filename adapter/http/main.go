package http

import (
	"github.com/bianavic/financial-planning-system/actuator"
	"github.com/bianavic/financial-planning-system/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transaction", transaction.AddTransaction)

	http.HandleFunc("/health", actuator.Health)
	http.Handle("/metrics", promhttp.Handler())

	_ = http.ListenAndServe(":8888", nil)
}
