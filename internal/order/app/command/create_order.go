package command

import (
	"context"

	"github.com/looksaw/go-orderv2/common/decorator"
	order2pb "github.com/looksaw/go-orderv2/common/genproto/orderpb"
	"github.com/looksaw/go-orderv2/order/app/query"
	domain "github.com/looksaw/go-orderv2/order/domain/order"
	"github.com/sirupsen/logrus"
)


type CreateOrder struct {
	CustomerID string
	Items []*order2pb.ItemWithQuantity
}


type CreateOrderResult struct {
	OrderID string
}

type CreateOrderHandler decorator.CommandHandler[CreateOrder,*CreateOrderResult]



type createOrderHandler struct {
	orderRepo domain.Repository
	stockGRPC query.StockService
}


func NewCreateOrderHandler(
	orderRepo domain.Repository,
	stockGRPC query.StockService,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
)CreateOrderHandler{
	if orderRepo == nil {
		panic("orderRepo must be valid")
	}
	return decorator.ApplyCommandDecorators[CreateOrder,*CreateOrderResult](
		createOrderHandler{orderRepo: orderRepo, stockGRPC: stockGRPC},
		logger,
		metricsClient,
	)
}


func (c createOrderHandler)Handle(ctx context.Context, cmd CreateOrder)(*CreateOrderResult,error){
	err := c.stockGRPC.CheckIfItemInStock(ctx,cmd.Items)
	logrus.Info("createOrderHandler|| err from stockGRPC ",err)
	var stockResponse []*order2pb.Item
	for _ , item := range cmd.Items {
		stockResponse = append(stockResponse, &order2pb.Item{
			ID: item.ID,
			Quantity: item.Quantity,
		})
	}
	o ,err := c.orderRepo.Create(ctx,&domain.Order{
		CustomerID: cmd.CustomerID,
		Item: stockResponse,
	})
	if err != nil {
		return nil ,err
	}
	return &CreateOrderResult{OrderID: o.ID} , nil
}