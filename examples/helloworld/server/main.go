package main

import (
	"context"
	"fmt"

	//"google.golang.org/grpc"

	"github.com/willieso/baby-univ-biz-service/examples/helloworld/helloworld"
	pb "github.com/willieso/baby-univ-biz-service/examples/helloworld/helloworld"
	grpcSrv "github.com/willieso/baby-univ-biz-service/internal/server"
	"github.com/willieso/baby-univ-biz-service/pkg/app"
	eagle "github.com/willieso/baby-univ-biz-service/pkg/app"
	logger "github.com/willieso/baby-univ-biz-service/pkg/log"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if in.Name == "" {
		return nil, fmt.Errorf("invalid argument %s", in.Name)
	}
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %+v", in.Name)}, nil
}

func main() {
	cfg := &app.ServerConfig{
		Network:      "tcp",
		Addr:         ":9090",
		ReadTimeout:  200,
		WriteTimeout: 200,
	}

	grpcServer := grpcSrv.NewGRPCServer(cfg)
	srv := &server{}
	helloworld.RegisterGreeterServer(grpcServer, srv)

	// start app
	app := eagle.New(
		eagle.WithName("eagle"),
		eagle.WithVersion("v1.0.0"),
		eagle.WithLogger(logger.GetLogger()),
		eagle.WithServer(
			grpcServer,
		),
	)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
