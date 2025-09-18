package service

import (
	"context"

	grpcClient "github.com/looksaw/go-orderv2/common/client"
	"github.com/looksaw/go-orderv2/common/metrics"
	"github.com/looksaw/go-orderv2/order/adapters"
	"github.com/looksaw/go-orderv2/common/broker"
	"github.com/looksaw/go-orderv2/order/adapters/grpc"
	"github.com/looksaw/go-orderv2/order/app"
	"github.com/looksaw/go-orderv2/order/app/command"
	"github.com/looksaw/go-orderv2/order/app/query"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewApplication(ctx context.Context) (app.Application,func()) {
	stockClient , closeStockClient , err := grpcClient.NewStockGRPCClient(ctx)
	if err != nil {
		panic(err)
	}
	ch , closeCh , err := broker.Connect(
		viper.GetString("rabbit-mq.user"),
		viper.GetString("rabbit-mq.password"),
		viper.GetString("rabbit-mq.host"),
		viper.GetString("rabbit-mq.port"),
	)
	stockGRPC := grpc.NewStockGRPC(stockClient)
	return newApplication(ctx,stockGRPC,ch) , func() {
		_ = closeStockClient()
		_ = ch.Close()
		_ = closeCh()
	}
}


func newApplication(_ context.Context , stockGRPC query.StockService,ch *amqp.Channel) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricsClient := metrics.TODOMetrics{}
	return app.Application{
		Commands: app.Commands{
			CreateOrder: command.NewCreateOrderHandler(orderRepo,stockGRPC,ch,logger,metricsClient),
			UpdateOrder: command.NewUpdateOrderHandler(orderRepo,logger,metricsClient),
		},
		Queries: app.Queries{
			GetCustomerOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, metricsClient),
		},
	}
}
