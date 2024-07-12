package payload

type ProductCreate struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,min=1" `
	Stock int     `json:"stock" validate:"min=0"`
}

type ProductUpdate struct {
	Name  *string  `json:"name,omitempty"`  // Optional field
	Price *float64 `json:"price,omitempty"` // Optional field
	Stock *int     `json:"stock,omitempty"` // Optional field
}
