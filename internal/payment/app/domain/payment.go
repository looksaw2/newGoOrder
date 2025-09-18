package domain

import (
	"context"

	"github.com/looksaw/go-orderv2/common/genproto/orderpb"
)

type Processor interface {
	CreatePaymentLink(ctx context.Context,pb *orderpb.Order)(string ,error)
}