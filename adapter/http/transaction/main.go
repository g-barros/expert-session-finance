package transaction

import (
	"net/http"
	"github.com/g-barros/expert-session-finance/model/transaction"
	"github.com/g-barros/expert-session-finance/util"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

// GetTransactions lista as transações
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title: "Salário",
			Amount: 1200.0,
			Type: 0,
			CreatedAt: util.StringToTime("2020-09-15T13:49:50"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

// CreateATransaction cria uma transação
func CreateATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}

	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}