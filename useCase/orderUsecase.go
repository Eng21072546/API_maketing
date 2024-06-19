package useCase

import (
	"github.com/Eng21072546/API_maketing/entity"
)

type OrderUseCase interface {
	CreateOrder(order entity.Order) (entity.Order, []string)
}

type OrderUseCaseImpl struct {
	repo OrderRepository
}

func (o OrderUseCaseImpl) CreateOrder(order entity.Order) (entity.Order, []string) {
	return o.repo.SaveOrder(order)
}

func NewOrderUseCase(repo OrderRepository) OrderUseCase {
	return &OrderUseCaseImpl{repo: repo}
}
