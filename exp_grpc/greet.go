package main

import (
	"fmt"
	"research/exp_grpc/proto"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (*Server) SayHello(req *proto.HelloRequest, stream proto.Greeter_SayHelloServer) error {
	someone := req.Name

	stream.Send(&proto.HelloResponse{
		Message: fmt.Sprintf("Keep it up! %s", someone),
	})

	return nil
}
