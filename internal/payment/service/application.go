package service

import (
	"context"

	grpc_client "github.com/looksaw/go-orderv2/common/client"
	"github.com/looksaw/go-orderv2/common/metrics"
	"github.com/looksaw/go-orderv2/payment/adapters"
	"github.com/looksaw/go-orderv2/payment/app"
	"github.com/looksaw/go-orderv2/payment/app/command"
	"github.com/looksaw/go-orderv2/payment/app/domain"
	"github.com/looksaw/go-orderv2/payment/infrastructure/processor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewApplication(ctx context.Context)(app.Application ,func()){
	orderClient , closeOrderClient , err := grpc_client.NewOrderGRPCClient(ctx)
	if err != nil {
		panic(err)
	}
	orderGRPC := adapters.NewOrderGRPC(orderClient)
	stripeProcessor := processor.NewStrpeProcessor(viper.GetString("stripe-key"))
	return newApplication(ctx,orderGRPC,stripeProcessor),func() {
		_ = closeOrderClient()
	}
}

func newApplication(ctx context.Context,orderGRPC command.OrderService,processor domain.Processor) app.Application{
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricClient := metrics.TODOMetrics{}
	return app.Application{
		Command: app.Command{
			CreatePayment: command.NewCreatePaymentHandler(
				processor,
				orderGRPC,
				logger,
				metricClient,
			),
		},
	}
}
