package main

import (
	"context"
	"github.com/jon-wade/oriClient/cli"
	pb "github.com/jon-wade/oriServer/ori"
	"google.golang.org/grpc"
	"log"
	"time"
)

// TODO: pop into an external config
const (
	address     = "localhost:50051"
)

func main() {
	// set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("connection not closed: %v", err)
		}
	}()

	// create gRPC client
	c := pb.NewMathHelperClient(conn)

	// create a context for the connection
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cli.Init(ctx, c)
}