package main

import (
	"context"
	"github.com/jon-wade/oriClient/cli"
	"time"
)

func main() {
	// create a context for the gRPC connection
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	cli.Init(ctx)
}