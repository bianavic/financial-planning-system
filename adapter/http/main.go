package http

import (
	"github.com/bianavic/financial-planning-system/actuator"
	"github.com/bianavic/financial-planning-system/adapter/http/transaction"
	"net/http"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transaction", transaction.AddTransaction)

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8888", nil)
}
