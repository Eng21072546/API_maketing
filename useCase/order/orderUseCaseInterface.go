package order

import (
	"context"
	"github.com/Eng21072546/API_maketing/entity"
)

type OrderUseCase interface {
	//GetOrderTransaction(id string) (*entity.Transaction, error)
	PatchOrderStatus(ctx context.Context, id string) (*entity.Order, error)
	NewOrder(Ctx context.Context, order *entity.Order) (*entity.Order, []error)
}
