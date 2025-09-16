package main

import (
	"log"
	"github.com/looksaw/go-orderv2/common/broker"
	"github.com/looksaw/go-orderv2/common/config"
	"github.com/looksaw/go-orderv2/common/server"
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