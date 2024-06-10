package models

type Order struct {
	ID           int      `json: "id"`
	customerName string   `json: "customerName"`
	Status       string   `json: "status"`
	ProductList  struct{} `json: "productList"`
}
