package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/jon-wade/oriServer/ori"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"time"
)

// TODO: pop into an external config
const (
	address     = "localhost:50051"
)

func summation(ctx context.Context, c pb.MathHelperClient, first int64, last int64) {
	s, err := c.Summation(ctx, &pb.SummationInput{First: first, Last: last})
	if err != nil {
		log.Fatalf("summationError: %v", err)
	}

	log.Printf("summationResult: %d", s.GetResult())
}

func factorial(ctx context.Context, c pb.MathHelperClient, base int64) {
	f, err := c.Factorial(ctx, &pb.FactorialInput{Base: base})
	if err != nil {
		log.Fatalf("factorialError: %v", err)
	}

	log.Printf("factorialResult: %d", f.GetResult())
}

func cli(ctx context.Context, c pb.MathHelperClient) {
	summationCmd := flag.NewFlagSet("summation", flag.ExitOnError)
	factorialCmd := flag.NewFlagSet("factorial", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'summation' or 'factorial' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "summation":
		err := summationCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal("could not parse subcommand:", err)
		}

		summationArgs := summationCmd.Args()
		if len(summationArgs) < 2 {
			fmt.Println("summation requires two arguments")
			os.Exit(1)
		}

		first, err := strconv.ParseInt(summationArgs[0], 10, 64)
		if err != nil {
			fmt.Println("arguments must be integers")
			os.Exit(1)
		}

		last, err := strconv.ParseInt(summationArgs[1], 10, 64)
		if err != nil {
			fmt.Println("arguments must be integers")
			os.Exit(1)
		}

		summation(ctx, c, first, last)
	case "factorial":
		err := factorialCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal("could not parse subcommand:", err)
		}

		factorialArgs := factorialCmd.Args()
		if len(factorialArgs) < 1 {
			fmt.Println("factorial requires one argument")
			os.Exit(1)
		}

		base, err := strconv.ParseInt(factorialArgs[0], 10, 64)
		if err != nil {
			fmt.Println("argument must be an integer")
			os.Exit(1)
		}

		factorial(ctx, c, base)
	default:
		fmt.Println("expected 'summation' or 'factorial' subcommands")
		os.Exit(1)
	}
}

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

	cli(ctx, c)
}