package client_test

import (
	"context"
	"github.com/jon-wade/oriClient/client"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		address string
	}{
		{"localhost:50051"},
		{"helloworld"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for idx, testData := range tests {
		_, err := client.Connect(ctx, testData.address)
		switch idx {
		case 0:
			if err != nil {
				t.Errorf("Expected err=%v, got %s", nil, err.Error())
			}
		case 1:
			if err == nil {
				t.Errorf("Expected err!=%v, got %s", nil, err)
			}
		}
	}
}