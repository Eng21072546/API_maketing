package payload

type ProductCreate struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type ProductUpdate struct {
	Name  *string  `json:"name,omitempty"`  // Optional field
	Price *float64 `json:"price,omitempty"` // Optional field
	Stock *int     `json:"stock,omitempty"` // Optional field
}
