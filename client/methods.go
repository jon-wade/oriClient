package client

import (
	"context"
	pb "github.com/jon-wade/oriServer/ori"
)

func Summation(ctx context.Context, c pb.MathHelperClient, first int64, last int64) (int64, error) {
	s, err := c.Summation(ctx, &pb.SummationInput{First: first, Last: last})
	if err != nil {
		return 0, err
	}

	return s.GetResult(), nil
}

func Factorial(ctx context.Context, c pb.MathHelperClient, base int64) (int64, error) {
	f, err := c.Factorial(ctx, &pb.FactorialInput{Base: base})
	if err != nil {
		return 0, err
	}

	return f.GetResult(), nil
}