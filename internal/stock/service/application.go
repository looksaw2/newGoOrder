package service

import (
	"context"

	"github.com/looksaw/go-orderv2/common/metrics"
	"github.com/looksaw/go-orderv2/stock/adapters"
	"github.com/looksaw/go-orderv2/stock/app"
	"github.com/looksaw/go-orderv2/stock/app/query"
	"github.com/sirupsen/logrus"
)


func NewApplication(ctx context.Context) app.Application {
	stockRepo := adapters.NewMemoryStockRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.TODOMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			CheckIfItemInStock: query.NewCheckIfItemsInStockHandler(stockRepo,logger,metricsClient),
			GetItems: query.NewGetItemsHandler(stockRepo,logger,metricsClient),
		},
	}
}