package payload

type LogisticCost struct {
	Address string  `json:"address" validate:"required"`
	Cost    float64 `json:"cost" validate:"required,min=0"`
}
