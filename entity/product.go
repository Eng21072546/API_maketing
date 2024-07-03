package entity

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

type ProductUpdate struct {
	Name  *string
	Price *float64
	Stock *int
}

type Stock struct {
	ID         int
	Name       string
	Quantities int
}
