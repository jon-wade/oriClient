package cli

import (
	"context"
	"flag"
	"fmt"
	"github.com/jon-wade/oriClient/cli/validate"
	"github.com/jon-wade/oriClient/client"
	pb "github.com/jon-wade/oriServer/ori"
	"google.golang.org/grpc"
	"log"
)

// as these are default values and not sensitive, have not added further config complexity by extracting as env vars
const (
	DefaultPort = 50051
	DefaultHost = "localhost"
)

func Init(ctx context.Context) {
	host := flag.String("host", DefaultHost, "hostname of oriserver, e.g. localhost")
	port := flag.Int("port", DefaultPort, "port number of oriserver, e.g. 50051")
	method := flag.String("method", "", "math helper method, e.g. summation or factorial")

	flag.Parse()
	args := flag.Args()

	// establish connection to server
	fmt.Printf("connecting to %s:%d\n", *host, *port)
	conn, err := client.Connect(ctx, fmt.Sprintf("%s:%d", *host, *port), grpc.DialContext)
	if err != nil {
		log.Fatal(err.Error())
	}

	// create gRPC client
	c := pb.NewMathHelperClient(conn)

	switch *method {
	case "summation":
		first, last, err := validate.SummationInputs(args)
		if err != nil {
			log.Fatal(err.Error())
		}

		result, err := client.Summation(ctx, c, first, last)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Printf("summationResult: %d", result)
		break
	case "factorial":
		base, err := validate.FactorialInputs(args)
		if err != nil {
			log.Fatal(err.Error())
		}

		result, err := client.Factorial(ctx, c, base)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Printf("factorialResult: %d", result)
		break
	default:
		log.Fatal("expected -method flag to be set to 'summation' or 'factorial'")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("connection not closed: %v", err)
		}
	}()
}