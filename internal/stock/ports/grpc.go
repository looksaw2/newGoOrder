package ports

import (
	"context"

	pb "github.com/looksaw/go-orderv2/common/genproto/stockpb"
	"github.com/looksaw/go-orderv2/stock/app"
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
	panic("TODO")
}

func(g *GRPCServer)CheckIfItemsInStock(context.Context, *pb.CheckIfItemsInStockRequest) (*pb.CheckIfItemsInStockResponse, error){
	panic("TODO")
}