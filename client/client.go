package client

import (
	pb "github.com/jon-wade/oriServer/ori"
	"google.golang.org/grpc"
	"log"
)

func Connect(address string) (pb.MathHelperClient, *grpc.ClientConn) {
	// set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// create gRPC client
	c := pb.NewMathHelperClient(conn)

	return c, conn
}