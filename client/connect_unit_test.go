package client_test

import (
	"context"
	"errors"
	"github.com/jon-wade/oriClient/client"
	"google.golang.org/grpc"
	"testing"
	"time"
)

// here we create a mock implementation of the DialContext function to avoid having to establish a connection
var callIdx = 0
func mockDialContext(_ context.Context, _ string, _ ...grpc.DialOption) (*grpc.ClientConn, error) {
	callIdx++
	var conn *grpc.ClientConn
	var err error
	switch callIdx {
	case 1:
		conn = &grpc.ClientConn{}
		break
	case 2:
		conn = nil
		err = errors.New("mock connection failure")
		break
	}

	return conn, err
}

func TestConnectUnit(t *testing.T) {
	tests := []struct {
		address string
	}{
		{"localhost:50051"},
		{"helloworld"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for idx, testData := range tests {
		_, err := client.Connect(ctx, testData.address, mockDialContext)
		switch idx {
		case 0:
			if err != nil {
				t.Errorf("Expected err=%v, got %s", nil, err.Error())
			}
		case 1:
			if err == nil {
				t.Errorf("Expected err!=%v, got %s", nil, err)
			}
		}
	}
}