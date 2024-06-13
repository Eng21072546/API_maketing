package models

type Order struct {
	ID           int            `json: "id"`
	CustomerName string         `json: "customername"`
	Address      string         `json: "address"`
	Status       Status         `json: "status"`
	ProductList  []ProductOrder `json: "productList"`
}

type ProductOrder struct {
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}

func CalculateOrderPrice() {

}
