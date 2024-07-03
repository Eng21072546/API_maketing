package useCase

import (
	"context"
	"errors"
	"fmt"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/google/uuid"
	"time"
)

type OrderUseCase interface {
	//GetOrderTransaction(id string) (*entity.Transaction, error)
	PatchOrderStatus(ctx context.Context, id string) (*entity.Order, error)

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
	transaction, err := o.transactionRepo.FindTransaction(ctx, order.TransactionId)
	if err != nil {
		errList = append(errList, err)
	}
	order.Transaction = transaction

	for _, item := range order.Transaction.ProductOrder {
		err := o.productRepo.CheckStock(ctx, item.ProductID, item.Quantity)
		if err != nil {
			errList = append(errList, err)
		}
		checked = append(checked, item)
	}

	if len(errList) > 0 { // If the errList has error stop and not reduce stock
		return nil, errList
	}

	err = o.productRepo.DecreaseStock(ctx, order.Transaction.ProductOrder)
	if err != nil {
		errList = append(errList, err)
	}
	orderTotalPrice += transaction.TotalPrice

	if len(errList) > 0 {
		return nil, errList
	}
	order.Status = entity.New
	order.ID = uuid.New().String()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	_, err = o.orderRepo.InsertOrder(ctx, collection.NewOrder(order))
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
	transaction, err := o.transactionRepo.FindTransaction(ctx, order.TransactionId)
	if err != nil {
		return nil, errors.New("transaction not found")
	}
	order.Transaction = transaction

	currStatus := order.Status
	newStatus := entity.UpStatus(currStatus)
	err = o.orderRepo.UpdateOrderStatus(ctx, id, newStatus) //update status
	if err != nil {
		return nil, errors.New("order status update failed")
	}
	order.Status = newStatus

	fmt.Println("Order ID %d confrim ", order.ID, " Status ", order.Status, "--> ", newStatus)

	return order, nil
}
