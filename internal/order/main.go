package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/looksaw/go-orderv2/common/broker"
	"github.com/looksaw/go-orderv2/common/config"
	"github.com/looksaw/go-orderv2/common/discovery"
	"github.com/looksaw/go-orderv2/common/genproto/orderpb"
	"github.com/looksaw/go-orderv2/common/server"
	"github.com/looksaw/go-orderv2/order/ports"
	"github.com/looksaw/go-orderv2/order/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

//初始化init
//读取配置
func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

//
func main(){
	serviceName := viper.GetString("order.service-name")
	ctx , cancel := context.WithCancel(context.Background())
	defer cancel()
	application,cleanup := service.NewApplication(ctx)
	defer cleanup()
	deregisterFunc ,err := discovery.RegisterToConsul(ctx,serviceName)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func ()  {
		_ = deregisterFunc()
	}()
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

	go server.RunGRPCServer(serviceName,func(server *grpc.Server) {
		svc := ports.NewGRPCServer(application)
		orderpb.RegisterOrderServiceServer(server,svc)		
	})
	server.RunHTTPServer(serviceName,func(router *gin.Engine) {
		ports.RegisterHandlersWithOptions(router,HTTPServer{
			app: application,
		},ports.GinServerOptions{
			BaseURL: "/api",
			Middlewares: nil,
			ErrorHandler: nil,
		})
	})
}
