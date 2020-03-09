package client_test

import (
	"context"
	"fmt"
	"github.com/jon-wade/oriClient/cli"
	"github.com/jon-wade/oriClient/client"
	pb "github.com/jon-wade/oriServer/ori"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"os"
	"testing"
	"time"
)

func TestSummationE2E(t *testing.T) {
	tests := []struct {
		first int64
		last int64
		result int64
		err error
	}{
		{2, 2, 4, nil},
		{math.MaxInt64, 1, 0, status.Errorf(codes.OutOfRange, "summation result exceeds maximum integer value")},
		{-250, 250, 0, nil},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := client.Connect(ctx, fmt.Sprintf("%s:%d", cli.DefaultHost, cli.DefaultPort), grpc.DialContext)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	c := pb.NewMathHelperClient(conn)

	for _, testData := range tests {
		testName := fmt.Sprintf("first=%d,last=%d,result=%d",
			testData.first, testData.last, testData.result)
		t.Run(testName, func(t *testing.T) {
			result, err := client.Summation(ctx, c, testData.first, testData.last)
			if result != testData.result {
				t.Errorf("Expected result=%d, got %v", testData.result, result)
			}
			if err != nil && err.Error() != testData.err.Error() {
				t.Errorf("Expected err=%s, got %s", testData.err.Error(), err.Error())
			}
		})
	}
}