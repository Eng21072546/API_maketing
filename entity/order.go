package entity

import "time"

type Order struct {
	ID           string `json:"id"`
	CustomerName string
	Status       Status         `json:"status"`
	Transaction  []*Transaction `json:"transaction"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	Amount       int            `json:"amount"`
	Total        float64        `json:"total"`
}

func UpStatus(status Status) Status {
	var newStatus Status
	if status == New {
		newStatus = Paid
	} else if status == Paid {
		newStatus = Processing
	} else if status == Processing {
		newStatus = Done
	} else {
		newStatus = Done
	}
	return newStatus
}
