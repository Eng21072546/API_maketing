package entity

import (
	"errors"
)

var logisticPrice = map[string]float64{
	"domestic":      40.0,
	"international": 100.0,
}

func CheckAddress(address string) error {
	_, ok := logisticPrice[address]
	if !ok {
		return errors.New("Invalid address")
	}
	return nil
}

func LogisticCost(address string) (float64, error) {
	err := CheckAddress(address)
	if err != nil {
		return 0, err
	}
	return logisticPrice[address], nil
}
