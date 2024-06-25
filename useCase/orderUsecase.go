package useCase

import (
	"context"
	"errors"
	"fmt"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/payload"
	"github.com/google/uuid"
	"time"
)

type OrderUseCase interface {
	//GetOrderTransaction(id string) (*entity.Transaction, error)
	PatchOrderStatus(ctx context.Context, id string) (*entity.Order, error)
	NewOrderEntity(ctx context.Context, orderPayload payload.Order) (*entity.Order, []error)
	NewOrder(Ctx context.Context, order *entity.Order) (*entity.Order, []error)
}

type OrderUseCaseImpl struct {
	orderRepo       OrderRepository
	productRepo     ProductRepository
	transactionRepo TransactionRepository
}

func NewOrderUseCase(orderRepo OrderRepository, productRepo ProductRepository, transactionRepo TransactionRepository) OrderUseCase {
	return &OrderUseCaseImpl{orderRepo: orderRepo, productRepo: productRepo, transactionRepo: transactionRepo}
}

func (o *OrderUseCaseImpl) NewOrder(ctx context.Context, order *entity.Order) (*entity.Order, []error) {
	var errList []error
	var checked []entity.ProductOrder
	var orderTotalPrice float64
	for _, transaction := range order.Transaction {
		for _, item := range transaction.ProductOrder {
			err := o.productRepo.CheckStock(item.ProductID, item.Quantity)
			if err != nil {
				errList = append(errList, err)
			}
			checked = append(checked, item)
		}
	}
	if len(errList) > 0 { // If the errList has error stop and not reduce stock
		return nil, errList
	}

	for _, transaction := range order.Transaction {
		err := o.productRepo.DecreaseStock(transaction.ProductOrder)
		if err != nil {
			errList = append(errList, err)
		}
		orderTotalPrice += transaction.TotalPrice
	}

	if len(errList) > 0 {
		return nil, errList
	}
	order.Status = entity.New
	order.ID = uuid.New().String()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.Amount = len(order.Transaction)
	order.Total = orderTotalPrice
	_, err := o.orderRepo.InsertOrder(ctx, collection.NewOrder(order))
	if err != nil {
		errList = append(errList, err)
	}
	if len(errList) > 0 {
		return nil, errList
	}
	return order, nil
}

func (o *OrderUseCaseImpl) PatchOrderStatus(ctx context.Context, id string) (*entity.Order, error) {
	order, err := o.orderRepo.FindOrderById(ctx, id)
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
	err = o.orderRepo.UpdateOrderStatus(ctx, id, newStatus) //update status
	if err != nil {
		return nil, errors.New("order status update failed")
	}
	order, err = o.orderRepo.FindOrderById(ctx, id)
	if err != nil {
		return nil, errors.New("order not found")
	}
	fmt.Println("Order ID %d confrim ", order.ID, " Status --> ", order.Status)

	return order, nil
}

func (o OrderUseCaseImpl) NewOrderEntity(ctx context.Context, orderPayload payload.Order) (*entity.Order, []error) {
	// This function use for create New order entity from payload, it can check transaction its has in DB?
	order := &entity.Order{}
	var errList []error
	var transactionList []*entity.Transaction
	for _, id := range orderPayload.TransactionId {
		transaction, err := o.transactionRepo.FindTransaction(ctx, id)
		if err != nil {
			errList = append(errList, errors.New(fmt.Sprintf("transaction %d not found", id)))
		} else {
			transactionList = append(transactionList, transaction)
		}
	}
	if len(errList) > 0 {
		return nil, errList
	}
	order.Transaction = transactionList
	order.CustomerName = orderPayload.CustomerName
	return order, nil
}
