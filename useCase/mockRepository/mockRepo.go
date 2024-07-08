package mockRepository

func NewMockRepository() (*MockProductRepo, *MockTransactionRepo, *MockOrderRepo, *MockLogisticRepo) {
	return new(MockProductRepo), new(MockTransactionRepo), new(MockOrderRepo), new(MockLogisticRepo)
}
