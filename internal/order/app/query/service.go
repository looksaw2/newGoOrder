package query

import (
	"context"

	order2pb "github.com/looksaw/go-orderv2/common/genproto/orderpb"
)

type StockService interface {
	CheckIfItemInStock(ctx context.Context ,items []*order2pb.ItemWithQuantity)error
	GetItems(ctx context.Context,itemIDs []string)([]*order2pb.Item ,error)
}