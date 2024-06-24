package entity

type Order struct {
	ID          string      `json:"id"`
	Status      Status      `json:"status"`
	Transaction Transaction `json:"transaction"`
}
