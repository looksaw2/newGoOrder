package query

import (
	"context"

	order2pb "github.com/looksaw/go-orderv2/common/genproto/orderpb"
	"github.com/looksaw/go-orderv2/common/genproto/stockpb"
)

type StockService interface {
	CheckIfItemInStock(ctx context.Context ,items []*order2pb.ItemWithQuantity)(*stockpb.CheckIfItemsInStockResponse,error)
	GetItems(ctx context.Context,itemIDs []string)([]*order2pb.Item ,error)
}