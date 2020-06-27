package tor

import "net"

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
