package entity

import "time"

type Order struct {
	ID            string
	CustomerName  string
	Status        Status
	TransactionId string
	Transaction   *Transaction
	CreatedAt     time.Time
	UpdatedAt     time.Time
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
