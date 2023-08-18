package main

import "research/exp_grpc/proto"

type Server struct {
	proto.UnimplementedGreeterServer
}

func (*Server) SayHello(req *proto.HelloRequest)
