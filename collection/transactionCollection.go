package collection

import (
	"github.com/Eng21072546/API_maketing/entity"
	"time"
)

type Transaction struct {
	ID           string                `json:"id"`
	Address      string                `json:"address"`
	AccountName  string                `json:"account_name"`
	Amount       int                   `json:"amount"`
	TotalPrice   float64               `json:"total_price"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
	ProductOrder []entity.ProductOrder `json:"product_order"`
}

func NewTransaction(transactionEntity *entity.Transaction) *Transaction {
	//This function for change Transaction Entity to Collection to save into DB
	transaction := new(Transaction)
	transaction.ID = transactionEntity.ID
	transaction.Address = transactionEntity.Address
	transaction.AccountName = transactionEntity.AccountName
	transaction.Amount = transactionEntity.Amount
	transaction.TotalPrice = transactionEntity.TotalPrice
	transaction.CreatedAt = transactionEntity.CreatedAt
	transaction.UpdatedAt = transactionEntity.UpdatedAt
	transaction.ProductOrder = transactionEntity.ProductOrder
	return transaction
}
