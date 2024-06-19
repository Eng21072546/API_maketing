package useCase

import "github.com/Eng21072546/API_maketing/entity"

type OrderRepository interface {
	SaveOrder(order entity.Order) (entity.Order, []string)
}
