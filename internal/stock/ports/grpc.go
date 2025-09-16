package ports

import (
	"context"

	//"errors"

	"github.com/looksaw/go-orderv2/common/genproto/stockpb"
	pb "github.com/looksaw/go-orderv2/common/genproto/stockpb"
	"github.com/looksaw/go-orderv2/stock/app"
	"github.com/looksaw/go-orderv2/stock/app/query"
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


func(g *GRPCServer)GetItems(ctx context.Context,request *pb.GetItemsRequest) (*pb.GetItemsResponse, error){
	logrus.Info("Start to get into the stock.ports...")
	items , err := g.app.Queries.GetItems.Handle(ctx,query.GetItems{ItemIDs: request.ItemIDs})
	if err != nil {
		return nil , err
	}
	return &stockpb.GetItemsResponse{Items: items} , nil
}

func(g *GRPCServer)CheckIfItemsInStock(ctx context.Context,request *pb.CheckIfItemsInStockRequest) (*pb.CheckIfItemsInStockResponse, error){
	items ,err := g.app.Queries.CheckIfItemInStock.Handle(ctx,query.CheckIfItemsInStock{Items: request.Items})
	if err != nil {
		return nil , err
	}	
	return &stockpb.CheckIfItemsInStockResponse{
		InStock: 1,
		Items: items,
	},nil
}