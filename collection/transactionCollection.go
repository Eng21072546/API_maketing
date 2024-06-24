package collection

import (
	"github.com/Eng21072546/API_maketing/entity"
	"time"
)

type Transaction struct {
	ID           string                `json:"id"`
	Address      string                `json:"address"`
	AccountName  string                `json:"account_name"`
	Amount       int64                 `json:"amount"`
	TotalPrice   float32               `json:"total_price"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
	ProductOrder []entity.ProductOrder `json:"product_order"`
}
