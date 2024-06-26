package entity

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type ProductUpdate struct {
	Name  *string  `json:"name" validate:"required"`
	Price *float64 `json:"price" validate:"required"`
	Stock *int     `json:"stock" validate:"required"`
}

type Stock struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Quantities int    `json:"quantities"`
}
