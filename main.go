package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	pb "github.com/midorigreen/goRpc/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8888"
)

type server struct {
}

func (g *server) GetSample(ctx context.Context, in *pb.SampleReq) (*pb.SampleRes, error) {
	return &pb.SampleRes{Mes: "message"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGoRpcServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
