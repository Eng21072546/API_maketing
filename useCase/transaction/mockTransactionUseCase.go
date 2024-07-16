package transaction

import (
	"context"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/stretchr/testify/mock"
)

type MockTransactionUseCase struct {
	mock.Mock
}

func (t *MockTransactionUseCase) NewTransaction(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, []error) {
	args := t.Called(ctx, transaction)
	if args.Get(0) == nil {
		return args.Get(0).(*entity.Transaction), nil
	}
	return args.Get(0).(*entity.Transaction), args.Get(1).([]error)
}

func (t *MockTransactionUseCase) FindTransactionById(context context.Context, id string) (*entity.Transaction, error) {
	args := t.Called(context, id)
	if args.Get(0) == nil {
		return args.Get(0).(*entity.Transaction), nil
	}
	return args.Get(0).(*entity.Transaction), args.Get(1).(error)
}
