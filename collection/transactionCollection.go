package collection

import (
	"github.com/Eng21072546/API_maketing/entity"
	"time"
)

type Transaction struct {
	ID           string
	Address      string
	Amount       int
	TotalPrice   float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ProductOrder []entity.ProductOrder
}

func NewTransaction(transactionEntity *entity.Transaction) *Transaction {
	//This function for change Transaction Entity to Collection to save into DB
	transaction := new(Transaction)
	transaction.ID = transactionEntity.ID
	transaction.Address = transactionEntity.Address
	transaction.Amount = transactionEntity.Amount
	transaction.TotalPrice = transactionEntity.TotalPrice
	transaction.CreatedAt = transactionEntity.CreatedAt
	transaction.UpdatedAt = transactionEntity.UpdatedAt
	transaction.ProductOrder = transactionEntity.ProductOrder
	return transaction
}
