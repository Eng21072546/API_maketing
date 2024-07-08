package mockUseCase

func NewMockUseCase() (*MockProductUseCase, *MockTransactionUseCase, *MockOrderUseCase) {
	return new(MockProductUseCase), new(MockTransactionUseCase), new(MockOrderUseCase)
}
