package tor

import (
	"context"
	"fmt"
	"testing"
)

func TestIsExit(t *testing.T) {
	var isExit bool

	queryDNSExitList = func(_ context.Context, _ string) ([]string, error) {
		if isExit {
			return []string{"127.0.0.2"}, nil
		}
		return nil, fmt.Errorf("NXDOMAIN")
	}

	var tests = []struct {
		isExit bool
		input  string
		want   bool
	}{
		// Positive
		{true, "0.0.0.0", true},

		// Negative
		{false, "9.9.9.9", false},
	}

	for _, test := range tests {
		isExit = test.isExit

		if got, _ := IsExit(context.TODO(), test.input); got != test.want {
			t.Errorf("IsExit(%s) = %t", test.input, got)
		}
	}
}
