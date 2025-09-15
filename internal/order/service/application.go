package service

import (
	"context"

	"github.com/looksaw/go-orderv2/common/metrics"
	"github.com/looksaw/go-orderv2/order/adapters"
	"github.com/looksaw/go-orderv2/order/app"
	"github.com/looksaw/go-orderv2/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.TODOMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, metricsClient),
		},
	}
}
