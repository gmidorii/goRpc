package main

import (
	"google.golang.org/grpc"
	"log"

	pb "github.com/midorigreen/goRpc/rpc"
	"context"
)

const (
	server = "localhost:8888"
)

func main() {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed connect : %v", err)
	}
	defer conn.Close()
	c := pb.NewGoRpcClient(conn)

	req := pb.SampleReq{
		Id: 1,
		Name: "name",
		Juice: pb.Juice_soda,
	}
	r, err := c.GetSample(context.Background(), &req)
	if err != nil {
		log.Fatalf("Failed get sample : %v", err)
	}
	log.Printf("Print message : %s", r.Mes)
}
