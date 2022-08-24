package order

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type orderCommandMock struct {
	mock.Mock
}

func (m *orderCommandMock) AddNewOrder(ctx context.Context, request OrderMasterEntity, detail []OrderDetailEntity) (*OrderWrap, error) {
	ret := m.Called(ctx, request, detail)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).(*OrderWrap), ret.Error(1)
}
