package models

import "errors"

var logisticPrice = map[string]float64{
	"domestic":      40.0,
	"international": 100.0,
}

func CheckAddress(order Order) error {
	_, ok := logisticPrice[order.Address]
	if !ok {
		return errors.New("Invalid address")
	}
	return nil
}

func LogisticCost(order Order) (float64, error) {
	err := CheckAddress(order)
	if err != nil {
		return 0, err
	}
	return logisticPrice[order.Address], nil
}
