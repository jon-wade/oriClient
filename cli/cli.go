package cli

import (
	"context"
	"errors"
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

func validateSummationInputs(args []string) (int64, int64, error) {
	if len(args) < 2 {
		return 0, 0, errors.New("summation requires two arguments")
	}

	first, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return 0, 0, errors.New("arguments must be integers")
	}

	last, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return 0, 0, errors.New("arguments must be integers")
	}

	return first, last, nil
}

func validateFactorialInputs(args []string) (int64, error) {
	if len(args) < 1 {
		return 0, errors.New("factorial requires one argument")
	}

	base, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return 0, errors.New("argument must be an integer")
	}

	return base, nil
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
		first, last, err := validateSummationInputs(args)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		methods.Summation(ctx, c, first, last)
	case "factorial":
		base, err := validateFactorialInputs(args)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
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