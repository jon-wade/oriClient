package client

import (
	"context"
	pb "github.com/jon-wade/oriServer/ori"
	"log"
)

func Summation(ctx context.Context, c pb.MathHelperClient, first int64, last int64) (int64, error) {
	s, err := c.Summation(ctx, &pb.SummationInput{First: first, Last: last})
	if err != nil {
		return 0, err
	}

	log.Printf("summationResult: %d", s.GetResult())

	return s.GetResult(), nil
}

func Factorial(ctx context.Context, c pb.MathHelperClient, base int64) int64 {
	f, err := c.Factorial(ctx, &pb.FactorialInput{Base: base})
	if err != nil {
		log.Fatalf("factorialError: %v", err)
	}

	log.Printf("factorialResult: %d", f.GetResult())

	return f.GetResult()
}