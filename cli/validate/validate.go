package validate

import (
	"errors"
	"strconv"
)

func SummationInputs(args []string) (int64, int64, error) {
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

func FactorialInputs(args []string) (int64, error) {
	if len(args) < 1 {
		return 0, errors.New("factorial requires one argument")
	}

	base, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return 0, errors.New("argument must be an integer")
	}

	return base, nil
}
