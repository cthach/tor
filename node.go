package tor

import (
	"context"
	"net"
)

// IsExit returns nil if the IP address belongs to an exit node.
func IsExit(ctx context.Context, ip net.IP) error {
	// TODO
	return nil
}

// queryDNSExitList queries the Tor Project's DNS based node exit list and returns the response, if any.
func queryDNSExitList(ctx context.Context, src net.IP, dst net.IP, port uint16) ([]net.IPAddr, error) {
	// TODO
	return nil, nil
}

// reverseIP returns an IP with octets in reverse order.
func reverseIP(ip net.IP) net.IP {
	// TODO
	return nil
}
