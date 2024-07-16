package transaction

import (
	"context"
	"github.com/Eng21072546/API_maketing/collection"
	"github.com/Eng21072546/API_maketing/entity"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepo struct {
	mock.Mock
}

func (t *MockTransactionRepo) InsertTransaction(ctx context.Context, transaction *collection.Transaction) (*entity.Transaction, error) {
	args := t.Called(ctx, transaction)
	return args.Get(0).(*entity.Transaction), args.Error(1)
}

func (t *MockTransactionRepo) FindTransaction(ctx context.Context, id string) (*entity.Transaction, error) {
	args := t.Called(ctx, id)
	return args.Get(0).(*entity.Transaction), args.Error(1)
}
