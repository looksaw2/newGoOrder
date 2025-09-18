package main

import (
	"context"
	"log"

	"github.com/looksaw/go-orderv2/common/broker"
	"github.com/looksaw/go-orderv2/common/config"
	"github.com/looksaw/go-orderv2/common/server"
	"github.com/looksaw/go-orderv2/payment/infrastructure/consumer"
	"github.com/looksaw/go-orderv2/payment/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)
func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main(){
	serverName := viper.GetString("payment.service-name")
	serverType := viper.GetString("payment.server-to-run")
	ctx ,cancel := context.WithCancel(context.Background())
	defer cancel()
	application , cleanup := service.NewApplication(ctx)
	defer cleanup()

	ch , closeCh , err := broker.Connect(
		viper.GetString("rabbit-mq.user"),
		viper.GetString("rabbit-mq.password"),
		viper.GetString("rabbit-mq.host"),
		viper.GetString("rabbit-mq.port"),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		_ = ch.Close()
		_ = closeCh()
	}()
	//处理rabbitmq中积累的消息
	go consumer.NewConsumer(application).Listen(ch)
	paymentHandler := NewPaymentHandler()
	switch serverType {
	case "http":
		server.RunHTTPServer(serverName,paymentHandler.RegisterRoutes)
	case "grpc":
		panic("unsupported type")
	default:
		panic("unsupported type")
	}
}