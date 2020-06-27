package tor

import (
	"context"
	"net"
)

// IsExit returns nil if the IP address belongs to an exit node.
func IsExit(ctx context.Context, ip string) error {
	// TODO
	return nil
}

// queryDNSExitList queries the Tor Project's DNS based node exit list and returns the response, if any.
func queryDNSExitList(ctx context.Context, src string, dst string, port uint16) ([]string, error) {
	// TODO
	return nil, nil
}

// isIP returns true if the string is a well-formed IP address.
func isIP(addr string) bool {
	if ip := net.ParseIP(addr); ip.To4() == nil && ip.To16() == nil {
		return false
	}
	return true
}

// reverse returns the reverse string.
func reverse(s string) string {
	runes := []rune(s)

	if len(runes) == 0 {
		return ""
	}

	for i, j := 0, len(runes)-1; i != j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
