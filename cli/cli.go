package cli

import (
	"context"
	"flag"
	"fmt"
	"github.com/jon-wade/oriClient/methods"
	pb "github.com/jon-wade/oriServer/ori"
	"log"
	"os"
	"strconv"
)

func Init(ctx context.Context, c pb.MathHelperClient) {
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

		methods.Summation(ctx, c, first, last)
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

		methods.Factorial(ctx, c, base)
	default:
		fmt.Println("expected 'summation' or 'factorial' subcommands")
		os.Exit(1)
	}
}