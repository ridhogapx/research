package main

import (
	"fmt"
	"net"
	"research/exp_grpc/proto"

	"google.golang.org/grpc"
)

var addr = "0.0.0.0:50051"

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(req *proto.HelloRequest, stream proto.Greeter_SayHelloServer) error {
	someone := req.Name

	err := stream.Send(&proto.HelloResponse{
		Message: fmt.Sprintf("Keep it up! %s", someone),
	})

	if err != nil {
		panic(err)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on", addr)

	s := grpc.NewServer()

	proto.RegisterGreeterServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
