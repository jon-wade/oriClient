package client_test

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"testing"
)

type mockMathHelperClient struct {}

func (c mockMathHelperClient) Summation() (int64, error) {
	callIdx++
	var result int64
	var err error
	switch callIdx {
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
		testName := fmt.Sprintf("first=%d,last=%d,result=%d",
			testData.first, testData.last, testData.result)
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