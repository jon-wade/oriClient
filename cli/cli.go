package cli

import (
	"context"
	"flag"
	"fmt"
	"github.com/jon-wade/oriClient/client"
	"github.com/jon-wade/oriClient/methods"
	"log"
	"os"
	"strconv"
)

const (
	defaultPort = 50051
	defaultHost = "localhost"
)

func validateSummationInputs(args []string) (int64, int64) {
	if len(args) < 2 {
		fmt.Println("summation requires two arguments")
		os.Exit(1)
	}

	first, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Println("arguments must be integers")
		os.Exit(1)
	}

	last, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		fmt.Println("arguments must be integers")
		os.Exit(1)
	}

	return first, last
}

func validateFactorialInputs(args []string) int64 {
	if len(args) < 1 {
		fmt.Println("factorial requires one argument")
		os.Exit(1)
	}

	base, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Println("argument must be an integer")
		os.Exit(1)
	}

	return base
}

func Init(ctx context.Context) {
	host := flag.String("host", defaultHost, "hostname of oriserver, e.g. localhost")
	port := flag.Int("port", defaultPort, "port number of oriserver, e.g. 50051")
	method := flag.String("method", "", "math helper method, e.g. summation or factorial")

	flag.Parse()
	args := flag.Args()

	// establish connection to server
	c, conn := client.Connect(fmt.Sprintf("%s:%d", *host, *port))
	fmt.Printf("connecting to %s:%d\n", *host, *port)

	switch *method {
	case "summation":
		first, last := validateSummationInputs(args)
		methods.Summation(ctx, c, first, last)
	case "factorial":
		base := validateFactorialInputs(args)
		methods.Factorial(ctx, c, base)
	default:
		fmt.Println("expected -method flag to be set to 'summation' or 'factorial'")
		os.Exit(1)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("connection not closed: %v", err)
		}
	}()
}