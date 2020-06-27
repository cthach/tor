package tor

import "context"

// Resolver is a DNS resolver.
type Resolver interface {
	// LookupHost returns a list of addresses for a host.
	LookupHost(ctx context.Context, host string) (addrs []string, err error)
}
