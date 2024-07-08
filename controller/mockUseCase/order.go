package mockUseCase

import (
	"context"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/stretchr/testify/mock"
)

type MockOrderUseCase struct {
	mock.Mock
}

func (o *MockOrderUseCase) PatchOrderStatus(ctx context.Context, id string) (*entity.Order, error) {
	args := o.Called(ctx, id)
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (o *MockOrderUseCase) NewOrder(Ctx context.Context, order *entity.Order) (*entity.Order, []error) {
	args := o.Called(Ctx, order)
	if args.Get(0) == nil {
		return args.Get(0).(*entity.Order), nil
	}
	return args.Get(0).(*entity.Order), args.Get(1).([]error)
}
