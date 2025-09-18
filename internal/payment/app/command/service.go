package command

import (
	"context"

	"github.com/looksaw/go-orderv2/common/genproto/orderpb"
)

type OrderService interface {
	UpdateOrder(ctx context.Context,order *orderpb.Order) error
	
}