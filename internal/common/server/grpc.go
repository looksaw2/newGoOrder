package server

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_tags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)
func init(){
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel)
	grpc_logrus.ReplaceGrpcLogger(logrus.NewEntry(logger))
}

func RunGRPCServer(serviceName string,registerServer func(server *grpc.Server)){
	addr := viper.Sub(serviceName).GetString("grpc-addr")
	if addr == ""{
		addr = viper.GetString("fallback-grpc-addr")
	}
	RunGRPCServerOnAddr(addr,registerServer)
}

func RunGRPCServerOnAddr(addr string,registerServer func(server *grpc.Server)){
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			grpc_tags.UnaryServerInterceptor(grpc_tags.WithFieldExtractor(grpc_tags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
		),
		grpc.ChainStreamInterceptor(
			grpc_tags.StreamServerInterceptor(grpc_tags.WithFieldExtractor(grpc_tags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry),

		),
	)
	registerServer(grpcServer)
	lis, err := net.Listen("tcp",addr)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Info("Starting gRPC server. Listening : %s",addr)
	if err := grpcServer.Serve(lis);  err != nil {
		logrus.Panic(err)
	}

}