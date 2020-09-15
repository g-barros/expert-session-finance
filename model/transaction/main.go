package transaction

import "time"

//Transaction estrutura uma transação
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

//Transactions uma coleção de transações
type Transactions []Transaction
