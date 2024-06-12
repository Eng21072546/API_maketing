package models

type Order struct {
	ID           int            `json: "id"`
	CustomerName string         `json: "customername"`
	Status       Status         `json: "status"`
	ProductList  []ProductOrder `json: "productList"`
}

type ProductOrder struct {
	ProductID int `json:"productId"` // Renamed for clarity
	Quantity  int `json:"quantity"`
}

//StatusOrder := ["New",""]
