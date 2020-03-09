package client

import (
	"context"
	"google.golang.org/grpc"
)

func Connect(ctx context.Context, address string) (*grpc.ClientConn, error) {
	// set up a connection to the server
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return conn, nil
}