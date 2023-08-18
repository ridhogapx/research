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
