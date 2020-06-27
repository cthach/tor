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
		if got := tor.IsIP(test.input); got != test.want {
			t.Errorf("isIP(%s)", test.input)
		}
	}
}

func TestReverse(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"hello", "olleh"},
	}

	for _, test := range tests {
		if got := tor.Reverse(test.input); got != test.want {
			t.Errorf("reverse(%s) = %s", test.input, got)
		}
	}
}
