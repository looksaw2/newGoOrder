package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/looksaw/go-orderv2/common/config"
	"github.com/looksaw/go-orderv2/common/genproto/orderpb"
	"github.com/looksaw/go-orderv2/common/server"
	"github.com/looksaw/go-orderv2/order/ports"
	"github.com/looksaw/go-orderv2/order/service"
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
	application := service.NewApplication(ctx)
	go server.RunGRPCServer(serviceName,func(server *grpc.Server) {
		svc := ports.NewGRPCServer(application)
		order2pb.RegisterOrderServiceServer(server,svc)		
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
