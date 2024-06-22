package useCase

import (
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/google/uuid"
	"strconv"
)

type OrderUseCase interface {
	CalculateOrderPrice(order entity.Order) (entity.Order, []string)
}

type OrderUseCaseImpl struct {
	orderRepo   OrderRepository
	productRepo ProductRepository
}

func NewOrderUseCase(orderRepo OrderRepository, productRepo ProductRepository) OrderUseCase {
	return &OrderUseCaseImpl{orderRepo: orderRepo, productRepo: productRepo}
}

func (o OrderUseCaseImpl) CalculateOrderPrice(order entity.Order) (entity.Order, []string) {
	var errList []string
	err := entity.CheckAddress(order)
	if err != nil {
		errList = append(errList, err.Error())
	}
	for _, productOrder := range order.ProductList {
		_, err = o.productRepo.FindProductById(productOrder.ProductID)
		if err != nil {
			errList = append(errList, err.Error())
		}
	}
	if len(errList) > 0 {
		return order, errList
	}
	order.ID = uuid.New().String()
	order.TotalPrice = o.CalculatePrice(order)
	_, err = o.orderRepo.InsertOrder(order)

	if err != nil {
		errList = append(errList, err.Error())
	}
	return order, errList
}

func (o OrderUseCaseImpl) CalculatePrice(order entity.Order) float64 {
	var totalPrice float64
	var bill []string
	logisticPrice, _ := entity.LogisticCost(order)
	totalPrice += logisticPrice
	for _, productOrder := range order.ProductList {
		product, _ := o.productRepo.FindProductById(productOrder.ProductID)
		productPrice := product.Price * float64(productOrder.Quantity)
		bill = append(bill, product.Name, " ", strconv.FormatFloat(product.Price, 'f', 2, 64), "  ", string(productOrder.Quantity), " ", strconv.FormatFloat(productPrice, 'f', 2, 64))
		totalPrice += productPrice
	}
	return totalPrice
}
