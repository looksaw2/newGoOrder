package ports

import (
	"context"
	"errors"

	pb "github.com/looksaw/go-orderv2/common/genproto/stockpb"
	"github.com/looksaw/go-orderv2/stock/app"
	"github.com/sirupsen/logrus"
)


type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{
		app: app,
	}
}


func(g *GRPCServer)GetItems(context.Context, *pb.GetItemsRequest) (*pb.GetItemsResponse, error){
	logrus.Info("rpc_request_in, stock.GetItems")
	defer func ()  {
		logrus.Info("rpc_request_out, stock.GetItems")
	}()
	return nil ,errors.New("fake err")
}

func(g *GRPCServer)CheckIfItemsInStock(context.Context, *pb.CheckIfItemsInStockRequest) (*pb.CheckIfItemsInStockResponse, error){
	logrus.Info("rpc_request_in, stock.CheckIfItemsInStock")
	defer func() {
		logrus.Info("rpc_request_out, stock.CheckIfItemsInStock")
	}()
	return nil , errors.New("fake err")
}