package grpc

import (
	"context"

	order2pb "github.com/looksaw/go-orderv2/common/genproto/orderpb"
	"github.com/looksaw/go-orderv2/common/genproto/stockpb"
	"github.com/sirupsen/logrus"
)

type StockGRPC struct {
	client stockpb.StockServiceClient
}

func NewStockGRPC(client stockpb.StockServiceClient) *StockGRPC {
	return &StockGRPC{
		client: client,
	}
}


func (s StockGRPC)CheckIfItemInStock(ctx context.Context ,items []*order2pb.ItemWithQuantity)error{
	resp , err := s.client.CheckIfItemsInStock(ctx,&stockpb.CheckIfItemsInStockRequest{
		Items: items,
	})
	logrus.Infof("stock_grpc response ",resp)
	return err
}

func (s StockGRPC)GetItems(ctx context.Context,itemIDs []string)([]*order2pb.Item ,error){
	resp , err := s.client.GetItems(ctx,&stockpb.GetItemsRequest{ItemIDs: itemIDs})
	if err != nil {
		return nil ,err
	}
	return resp.Items , nil
}