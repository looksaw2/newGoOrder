package ports

import (
	"context"

	pb "github.com/looksaw/go-orderv2/common/genproto/orderpb"
	"github.com/looksaw/go-orderv2/order/app"
	"google.golang.org/protobuf/types/known/emptypb"
)


type GRPCServer struct {}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{}
}

func(g *GRPCServer)CreateOrder(context.Context, *pb.CreateOrderRequest) (*emptypb.Empty, error){
	panic("Todo")
}

func(g *GRPCServer)GetOrder(context.Context, *pb.GetOrderRequest) (*pb.Order, error){
	panic("Todo")
}
func(g *GRPCServer)UpdateOrder(context.Context, *pb.Order) (*emptypb.Empty, error){
	panic("Todo")
}
