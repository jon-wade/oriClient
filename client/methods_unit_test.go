package client_test

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"testing"
)

var summationCallIdx = 0
var factorialCallIdx = 0

// here were mocking out the MathHelperClient by implementing the methods of the MathHelperClient interface
type mockMathHelperClient struct {}

func (c mockMathHelperClient) Summation() (int64, error) {
	summationCallIdx++
	var result int64
	var err error
	switch summationCallIdx {
	case 1:
		result = 4
		break
	case 2:
		result = 0
		err = status.Errorf(codes.OutOfRange, "summation result exceeds maximum integer value")
		break
	case 3:
		result = 0
		break
	}

	return result, err
}

func (c mockMathHelperClient) Factorial() (int64, error) {
	factorialCallIdx++
	var result int64
	var err error
	switch factorialCallIdx {
	case 1:
		result = 120
		break
	case 2:
		result = 0
		err = status.Errorf(codes.OutOfRange, "factorial result exceeds maximum integer value")
		break
	}

	return result, err
}

func TestSummationUnit(t *testing.T) {
	tests := []struct {
		first int64
		last int64
		result int64
		err error
	}{
		{2, 2, 4, nil},
		{math.MaxInt64, 1, 0, status.Errorf(codes.OutOfRange, "summation result exceeds maximum integer value")},
		{-250, 250, 0, nil},
	}

	for _, testData := range tests {
		testName := fmt.Sprintf("first=%d,last=%d,result=%d,err=%s",
			testData.first, testData.last, testData.result, testData.err)
		t.Run(testName, func(t *testing.T) {
			result, err := mockMathHelperClient{}.Summation()
			if result != testData.result {
				t.Errorf("Expected result=%d, got %v", testData.result, result)
			}
			if err != nil && err.Error() != testData.err.Error() {
				t.Errorf("Expected err=%s, got %s", testData.err.Error(), err.Error())
			}
		})
	}
}

func TestFactorialUnit(t *testing.T) {
	tests := []struct {
		base int64
		result int64
		err error
	}{
		{5, 120, nil},
		{21, 0, status.Errorf(codes.OutOfRange, "factorial result exceeds maximum integer value")},
	}

	for _, testData := range tests {
		testName := fmt.Sprintf("base=%d,result=%d,err=%s",
			testData.base, testData.result, testData.err)
		t.Run(testName, func(t *testing.T) {
			result, err := mockMathHelperClient{}.Factorial()
			if result != testData.result {
				t.Errorf("Expected result=%d, got %v", testData.result, result)
			}
			if err != nil && err.Error() != testData.err.Error() {
				t.Errorf("Expected err=%s, got %s", testData.err.Error(), err.Error())
			}
		})
	}
}