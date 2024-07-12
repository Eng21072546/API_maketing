package payload

type Order struct {
	CustomerName  string `json:"customerName" validate:"required"`
	TransactionId string `json:"transactionId" validate:"required"`
}
