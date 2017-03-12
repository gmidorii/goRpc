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

	r, err := c.GetSample(context.Background(), &pb.SampleReq{Id: 1, Name: "hoge"})
	if err != nil {
		log.Fatalf("Failed get sample : %v", err)
	}
	log.Printf("Print message : %s", r.Mes)
}
