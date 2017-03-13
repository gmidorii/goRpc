package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	pb "github.com/midorigreen/goRpc/rpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8888"
)

type server struct {
}

func (g *server) GetSample(ctx context.Context, in *pb.SampleReq) (*pb.SampleRes, error) {
	logrus.WithFields(logrus.Fields{
		"id" : in.Id,
		"name": in.Name,
	}).Info("request parameter")
	return &pb.SampleRes{Mes: "Hi! " + in.Name}, nil
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
