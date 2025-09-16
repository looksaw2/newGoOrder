package command

import (
	"context"

	"github.com/looksaw/go-orderv2/common/decorator"
	domain "github.com/looksaw/go-orderv2/order/domain/order"
	"github.com/sirupsen/logrus"
)

type UpdateOrder struct {
	Order *domain.Order
	UpdateFn func(context.Context , *domain.Order)(*domain.Order ,error)
}

type UpdateOrderHandler decorator.CommandHandler[UpdateOrder,any]

type updateOrderHandler struct {
	orderRepo domain.Repository
}


func NewUpdateOrderHandler(
	orderRepo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
)UpdateOrderHandler{
	if orderRepo == nil {
		panic("orderRepo must be valid")
	}
	return decorator.ApplyCommandDecorators[UpdateOrder,any](
		updateOrderHandler{orderRepo: orderRepo},
		logger,
		metricsClient,
	)
}


func (c updateOrderHandler)Handle(ctx context.Context, cmd UpdateOrder)(any,error){
	if cmd.UpdateFn == nil {
		logrus.Warnf("updateOrderHandler got nil UpdateFn, orderID=%v",cmd.Order.ID)
		cmd.UpdateFn = func(ctx context.Context, o *domain.Order) (*domain.Order, error) {
			return o , nil 
		}
	}
	err := c.orderRepo.Update(ctx,cmd.Order,cmd.UpdateFn)
	if err != nil {
		return nil ,err
	}
	return nil ,nil
}