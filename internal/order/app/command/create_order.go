package command

import (
	"context"
	"errors"

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
	logrus.Info("<Handle> start to order valid")
	validItems , err := c.validate(ctx,cmd.Items)
	if err != nil {
		return nil ,err
	}
	logrus.Info("<Handle> start to orderRepo Create")
	o ,err := c.orderRepo.Create(ctx,&domain.Order{
		CustomerID: cmd.CustomerID,
		Item: validItems,
	})
	if err != nil {
		return nil ,err
	}
	logrus.Info("finish ... No Error")
	return &CreateOrderResult{OrderID: o.ID} , nil
}

func (c createOrderHandler)validate(ctx context.Context,items []*order2pb.ItemWithQuantity)([]*order2pb.Item,error){
	if len(items) == 0 {
		logrus.Info("len(items) == 0")
		return nil , errors.New("must have at least one item")
	}
	items = packItems(items)
	logrus.Info("checckIfItemInStock")
	resp , err := c.stockGRPC.CheckIfItemInStock(ctx,items)
	if err != nil {
		return nil , err
	}
	return resp.Items , nil
}

func packItems(items []*order2pb.ItemWithQuantity) []*order2pb.ItemWithQuantity {
	merged := make(map[string]int32)
	for _ , item := range items {
		merged[item.ID] += item.Quantity
	}
	var res []*order2pb.ItemWithQuantity
	for id , quantity := range merged {
		res = append(res, &order2pb.ItemWithQuantity{
			ID: id,
			Quantity: quantity,
		})
	}
	return  res 
}