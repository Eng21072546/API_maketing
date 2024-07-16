package order

import (
	"context"
	"errors"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/Eng21072546/API_maketing/repo/order"
	"github.com/Eng21072546/API_maketing/repo/product"
	"github.com/Eng21072546/API_maketing/repo/transaction"
)

type OrderUseCaseImpl struct {
	orderRepo       order.OrderRepository
	productRepo     product.ProductRepository
	transactionRepo transaction.TransactionRepository
}

func NewOrderUseCase(orderRepo order.OrderRepository, productRepo product.ProductRepository, transactionRepo transaction.TransactionRepository) OrderUseCase {
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
	result, err := o.orderRepo.InsertOrder(ctx, collection.NewOrder(order))
	if err != nil {
		errList = append(errList, err)
	}
	if len(errList) > 0 {
		return nil, errList
	}
	return result, nil
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
	newStatus := order.UpStatus(currStatus)
	_, err = o.orderRepo.UpdateOrderStatus(ctx, id, newStatus) //update status
	if err != nil {
		return nil, errors.New("order status update failed")
	}
	order.Status = newStatus
	orderUpdated, err := o.orderRepo.FindOrderById(ctx, order.ID)
	if err != nil {
		return nil, errors.New("can not find order")
	}
	if orderUpdated.Status != newStatus {
		return nil, errors.New("order status update failed")
	}

	return orderUpdated, nil
}
