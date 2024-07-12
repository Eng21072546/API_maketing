package payload

import "github.com/Eng21072546/API_maketing/entity"

type Transaction struct {
	AccountName  string                `json:"accountName"`
	Address      string                `json:"address" validate:"required"`
	ProductOrder []entity.ProductOrder `json:"productOrder" validate:"required"`
}
