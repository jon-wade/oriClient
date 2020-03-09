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
	"log"
	"math"
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
		log.Fatal(err.Error())
	}

	c := pb.NewMathHelperClient(conn)
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("connection not closed: %v", err)
		}
	}()

	for _, testData := range tests {
		testName := fmt.Sprintf("first=%d,last=%d,result=%d,err=%s",
			testData.first, testData.last, testData.result, testData.err)
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

func TestFactorialE2E(t *testing.T) {
	tests := []struct {
		base int64
		result int64
		err error
	}{
		{5, 120, nil},
		{21, 0, status.Errorf(codes.OutOfRange, "factorial result exceeds maximum integer value")},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := client.Connect(ctx, fmt.Sprintf("%s:%d", cli.DefaultHost, cli.DefaultPort), grpc.DialContext)
	if err != nil {
		log.Fatal(err.Error())
	}

	c := pb.NewMathHelperClient(conn)
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("connection not closed: %v", err)
		}
	}()

	for _, testData := range tests {
		testName := fmt.Sprintf("base=%d,result=%d,err=%s",
			testData.base, testData.result, testData.err)
		t.Run(testName, func(t *testing.T) {
			result, err := client.Factorial(ctx, c, testData.base)
			if result != testData.result {
				t.Errorf("Expected result=%d, got %v", testData.result, result)
			}
			if err != nil && err.Error() != testData.err.Error() {
				t.Errorf("Expected err=%s, got %s", testData.err.Error(), err.Error())
			}
		})
	}
}