package useCase

import (
	"errors"
	"fmt"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/google/uuid"
	"strconv"
)

type OrderUseCase interface {
	CalculateOrderPrice(order entity.Order) (entity.Order, []string)
	CreateOrder(transaction *entity.Transaction) (*entity.Order, []string)
	GetOrderTransaction(id string) (*entity.Transaction, error)
	PatchOrderStatus(id string) (*entity.Order, error)
}

type OrderUseCaseImpl struct {
	orderRepo   OrderRepository
	productRepo ProductRepository
}

func NewOrderUseCase(orderRepo OrderRepository, productRepo ProductRepository) OrderUseCase {
	return &OrderUseCaseImpl{orderRepo: orderRepo, productRepo: productRepo}
}

func (o *OrderUseCaseImpl) CalculateOrderPrice(order entity.Order) (entity.Order, []string) {
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

func (o *OrderUseCaseImpl) CalculatePrice(order entity.Order) float64 {
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

func (o *OrderUseCaseImpl) GetOrderTransaction(id string) (*entity.Transaction, error) {
	order, err := o.orderRepo.FindOrderById(id)
	if err != nil {
		return nil, err
	}
	transaction := entity.Transaction{ID: id, ProductOrder: order.ProductList}
	return &transaction, nil
}

func (o *OrderUseCaseImpl) CreateOrder(transaction *entity.Transaction) (*entity.Order, []string) {
	var errList []string
	var checked []entity.ProductOrder
	order, err := o.orderRepo.FindOrderById(transaction.ID)
	if err != nil {
		errList = append(errList, errors.New("transaction not found").Error())
		return nil, errList
	}
	for _, item := range transaction.ProductOrder {
		err := o.productRepo.CheckStock(item.ProductID, item.Quantity)
		if err != nil {
			errList = append(errList, err.Error())
		}
		checked = append(checked, item)
	}
	if len(errList) > 0 {
		return nil, errList
	}
	err = o.productRepo.DecreaseStock(transaction.ProductOrder)
	if err != nil {
		errList = append(errList, err.Error())
	}
	if len(errList) > 0 {
		return nil, errList
	}
	err = o.orderRepo.UpdateOrderStatus(order.ID, entity.New)
	if err != nil {
		errList = append(errList, err.Error())
	}
	order.Status = entity.New
	return order, nil
}

func (o *OrderUseCaseImpl) PatchOrderStatus(id string) (*entity.Order, error) {
	order, err := o.orderRepo.FindOrderById(id)
	if err != nil {
		return nil, errors.New("order not found")
	}
	currStatus := order.Status
	var newStatus entity.Status
	if currStatus == entity.New {
		newStatus = entity.Paid
	} else if currStatus == entity.Paid {
		newStatus = entity.Processing
	} else if currStatus == entity.Processing {
		newStatus = entity.Done
	} else {
		newStatus = entity.Done
	}
	err = o.orderRepo.UpdateOrderStatus(id, newStatus) //update status
	if err != nil {
		return nil, errors.New("order status update failed")
	}
	order, err = o.orderRepo.FindOrderById(id)
	if err != nil {
		return nil, errors.New("order not found")
	}
	fmt.Println("Order ID %d confrim ", order.ID, " Status --> ", order.Status)

	return order, nil
}
