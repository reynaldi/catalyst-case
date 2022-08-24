package order

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type orderQueryMock struct {
	mock.Mock
}

func (o *orderQueryMock) GetOrders(ctx context.Context) ([]OrderWrap, error) {
	ret := o.Called(ctx)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).([]OrderWrap), ret.Error(1)
}

func (o *orderQueryMock) GetOrder(ctx context.Context, orderId int) (*OrderWrap, error) {
	ret := o.Called(ctx, orderId)
	if ret.Get(0) == nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0).(*OrderWrap), ret.Error(1)
}
