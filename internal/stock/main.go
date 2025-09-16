package main

import (
	"context"
	"log"

	"github.com/looksaw/go-orderv2/common/config"
	"github.com/looksaw/go-orderv2/common/discovery"
	"github.com/looksaw/go-orderv2/common/genproto/stockpb"
	"github.com/looksaw/go-orderv2/common/server"
	"github.com/looksaw/go-orderv2/stock/ports"
	"github.com/looksaw/go-orderv2/stock/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)
func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}
func main(){
	serviceName := viper.Sub("stock").GetString("service-name")
	serviceType := viper.GetString("stock.server-to-run")
	ctx ,cancel := context.WithCancel(context.Background())
	defer cancel()
	application := service.NewApplication(ctx)
	deregisterFunc ,err := discovery.RegisterToConsul(ctx,serviceName)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func ()  {
		_ = deregisterFunc()
	}()
	switch serviceType {
	case "grpc":
		server.RunGRPCServer(serviceName,func(server *grpc.Server) {
			svc := ports.NewGRPCServer(application)
			stockpb.RegisterStockServiceServer(server,svc)	
	})
	default:
		panic("unexpected server type")	
	}	
}