package query

import (
	"context"

	"github.com/looksaw/go-orderv2/common/decorator"
	"github.com/looksaw/go-orderv2/common/genproto/orderpb"
	domain "github.com/looksaw/go-orderv2/stock/domain/stock"
	"github.com/sirupsen/logrus"
)

type CheckIfItemsInStock struct {
	Items []*order2pb.ItemWithQuantity
}

type CheckIfItemsInStockHandler decorator.QueryHandler[CheckIfItemsInStock, []*order2pb.Item]

type checkIfItemsInStockHandler struct {
	stockRepo domain.Repository
}

func NewCheckIfItemsInStockHandler(
	stockRepo domain.Repository,
	logger *logrus.Entry,
	metricClient decorator.MetricsClient,
) CheckIfItemsInStockHandler {
	if stockRepo == nil {
		panic("nil stockRepo")
	}
	return decorator.ApplyQueryDecorators[CheckIfItemsInStock, []*order2pb.Item](
		checkIfItemsInStockHandler{stockRepo: stockRepo},
		logger,
		metricClient,
	)
}

func (h checkIfItemsInStockHandler) Handle(ctx context.Context, query CheckIfItemsInStock) ([]*order2pb.Item, error) {
	var res []*order2pb.Item
	for _, i := range query.Items {
		res = append(res, &order2pb.Item{
			ID:       i.ID,
			Quantity: i.Quantity,
		})
	}
	return res ,nil
}