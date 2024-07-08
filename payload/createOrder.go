package payload

type Order struct {
	CustomerName  string `json:"customerName"`
	TransactionId string `json:"transactionId"`
}
