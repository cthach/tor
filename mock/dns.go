package mock

import (
	"context"

	"github.com/cthach/tor"
)

// Resolver is a mock DNS resolver.
type Resolver struct {
	Addrs []string
}

var _ tor.Resolver = &Resolver{}

// LookupHost always returns r.Addrs and a nil error.
func (r *Resolver) LookupHost(_ context.Context, _ string) ([]string, error) {
	return r.Addrs, nil
}
