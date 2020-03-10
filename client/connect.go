package client

import (
	"context"
	"google.golang.org/grpc"
)

// this is the type definition of grpc.DialContext
type DialCtx func(ctx context.Context, target string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error)

// dialCtx allows us to pass a mocked grpc dialing function into the Connect function to create unit tests
func Connect(ctx context.Context, address string, dialCtx DialCtx) (*grpc.ClientConn, error) {
	// set up a connection to the server
	conn, err := dialCtx(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
