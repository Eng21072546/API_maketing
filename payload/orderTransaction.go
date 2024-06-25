package payload

import "github.com/Eng21072546/API_maketing/entity"

type Transaction struct {
	AccountName  string                `json:"account_name"`
	Address      string                `json:"address"`
	ProductOrder []entity.ProductOrder `json:"product_order"`
}
