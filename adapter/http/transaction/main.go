package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/bianavic/financial-planning-system/model/transaction"
	"io"
	"net/http"
	"time"
)

func GetTransactions(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	//fmt.Fprintf(writer, "Hello, Stranger! Welcome to my Home Page")

	var layout = "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2024-02-29T15:45:00")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salary",
			Amount:    5888.33,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(writer).Encode(transactions)
}

func AddTransaction(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	var res = transaction.Transaction{}
	var body, _ = io.ReadAll(request.Body) // array de bytes

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
