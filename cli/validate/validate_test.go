package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidateSummationInputsUnit(t *testing.T) {
	tests := []struct {
		args  []string
		first int64
		last  int64
		err   error
	}{
		{[]string{"2", "2"}, 2, 2, nil},
		{[]string{"a", "2"}, 0, 0, errors.New("arguments must be integers")},
		{[]string{"", ""}, 0, 0, errors.New("arguments must be integers")},
		{[]string{"5"}, 0, 0, errors.New("summation requires two arguments")},
	}

	for _, testData := range tests {
		testName := fmt.Sprintf("args=%v,first=%d,last=%d,err=%s",
			testData.args, testData.first, testData.last, testData.err)
		t.Run(testName, func(t *testing.T) {
			first, last, err := SummationInputs(testData.args)
			if first != testData.first {
				t.Errorf("Expected first=%d, got %v", testData.first, first)
			}
			if last != testData.last {
				t.Errorf("Expected last=%d, got %v", testData.last, last)
			}
			if err != nil && err.Error() != testData.err.Error() {
				t.Errorf("Expected err=%s, got %s", testData.err.Error(), err.Error())
			}
		})
	}
}

func TestValidateFactorialInputsUnit(t *testing.T) {
	tests := []struct {
		args []string
		base int64
		err  error
	}{
		{[]string{"2"}, 2, nil},
		{[]string{"a"}, 0, errors.New("argument must be an integer")},
		{[]string{""}, 0, errors.New("argument must be an integer")},
		{[]string{}, 0, errors.New("factorial requires one argument")},
	}

	for _, testData := range tests {
		testName := fmt.Sprintf("args=%v,base=%d,err=%s",
			testData.args, testData.base, testData.err)
		t.Run(testName, func(t *testing.T) {
			base, err := FactorialInputs(testData.args)
			if base != testData.base {
				t.Errorf("Expected base=%d, got %d", testData.base, base)
			}
			if err != nil && err.Error() != testData.err.Error() {
				t.Errorf("Expected err=%s, got %s", testData.err.Error(), err.Error())
			}
		})
	}
}
