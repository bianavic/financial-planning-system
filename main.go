package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", getTransactions)

	http.ListenAndServe(":8888", nil)

}

type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
}

type Transactions []Transaction

func getTransactions(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	//fmt.Fprintf(writer, "Hello, Stranger! Welcome to my Home Page")

	var layout = "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2024-02-29T15:45:00")

	var transactions = Transactions{
		Transaction{
			Title:     "Salary",
			Amount:    5888.33,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(writer).Encode(transactions)
}
