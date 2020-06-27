package tor_test

import (
	"testing"

	"github.com/cthach/tor"
)

func TestIsIP(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		// Positive
		{"0.0.0.0", true},
		{"127.0.0.1", true},

		// Negative
		{"", false},
		{"IP? I hardly knew ye", false},
	}

	for _, test := range tests {
		if got := tor.IsIPv4(test.input); got != test.want {
			t.Errorf("isIP(%s)", test.input)
		}
	}
}

func TestReverseIP(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"0.0.0.0", "0.0.0.0"},
		{"1.2.3.4", "4.3.2.1"},
	}

	for _, test := range tests {
		if got := tor.ReverseIP(test.input); got != test.want {
			t.Errorf("reverseIP(%s) = %s", test.input, got)
		}
	}
}
