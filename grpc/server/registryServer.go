package server

import (
	"context"
	"fmt"
	pb "github.com/liurongdev/firstGoWeb/grpc/proto"
	"github.com/liurongdev/firstGoWeb/middleware/logger"
	"google.golang.org/grpc"
	"net"
)

type HelloServiceServer struct {
	pb.UnimplementedHelloServiceServer
}

// 已经被潜逃方法实现，不需要再实现
//func (s *HelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {
//	fmt.Println("mustEmbedUnimplementedHelloServiceServer")
//}

func (s *HelloServiceServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("grpc call HelloServiceServer.sayHello", req.Name)
	return &pb.HelloResponse{Message: "Hello, " + req.Name}, nil
}

func StartGRPC(grpcListener net.Listener) {
	server := grpc.NewServer()
	pb.RegisterHelloServiceServer(server, &HelloServiceServer{})
	if err := server.Serve(grpcListener); err != nil {
		logger.Error(err.Error())
	}
}
