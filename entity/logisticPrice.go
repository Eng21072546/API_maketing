package entity

import (
	"errors"
	"fmt"
	"strings"
)

var logisticPrice = map[string]float64{
	"domestic":      40.0,
	"international": 100.0,
}

func CheckAddress(address string) error {
	addressLow := strings.ToLower(address)
	_, ok := logisticPrice[addressLow]
	if !ok {
		return errors.New("invalid address")
	}
	return nil
}

func LogisticCost(address string) (float64, error) {
	addressLow := strings.ToLower(address)
	err := CheckAddress(address)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return logisticPrice[addressLow], nil
}
