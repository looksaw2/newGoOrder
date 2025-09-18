package ports

import (
	"context"

	pb "github.com/looksaw/go-orderv2/common/genproto/orderpb"
	"github.com/looksaw/go-orderv2/order/app"
	"github.com/looksaw/go-orderv2/order/app/command"
	"github.com/looksaw/go-orderv2/order/app/query"
	domain "github.com/looksaw/go-orderv2/order/domain/order"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)


type GRPCServer struct {
	app app.Application
}

func NewGRPCServer(app app.Application) *GRPCServer {
	return &GRPCServer{
		app: app,
	}
}

func(g GRPCServer)CreateOrder(ctx context.Context, request *pb.CreateOrderRequest) (*emptypb.Empty, error){
	_ , err := g.app.Commands.CreateOrder.Handle(ctx,command.CreateOrder{
		CustomerID: request.CustomerID,
		Items: request.Items,
	})
	if err != nil {
		return nil ,status.Error(codes.NotFound,err.Error())
	}
	return &emptypb.Empty{} , nil 
}

func(g GRPCServer)GetOrder(ctx context.Context, request *pb.GetOrderRequest) (*pb.Order, error){
	o ,err := g.app.Queries.GetCustomerOrder.Handle(ctx,
		query.GetCustomerOrder{
			CustomerID: request.CustomerID,
			OrderID: request.OrderID,
		},
	)
	if err != nil {
		return nil ,status.Error(codes.NotFound,err.Error())
	}
	return o.ToProto() ,nil
}
func(g GRPCServer)UpdateOrder(ctx context.Context,request  *pb.Order) (_ *emptypb.Empty,err error){
	logrus.Infof("order_grpc || request_in ||  request=%+v",request)
	order ,err := domain.NewOrder(
		request.ID,
		request.CustomerID,
		request.Status,
		request.PaymentLink,
		request.Items,
	)
	if err != nil {
		err = status.Error(codes.Internal,err.Error())
		return
	}
	_ ,err = g.app.Commands.UpdateOrder.Handle(ctx,command.UpdateOrder{
		Order: order,
		UpdateFn: func(ctx context.Context, o *domain.Order) (*domain.Order, error) {
			return order , nil
		},
	})
	if err != nil {
		return nil ,err
	}
	return 	
}
