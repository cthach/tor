package tor

import "net"

// isIPv4 returns true if the string is a well-formed IPv4 address.
func isIPv4(addr string) bool {
	if ip := net.ParseIP(addr); ip.To4() == nil {
		return false
	}
	return true
}

// reverseIP returns an IP address with reversed octets.
func reverseIP(addr string) string {
	ip := net.ParseIP(addr)
	if ip == nil {
		return ""
	}

	if ip.IsUnspecified() {
		return ip.String()
	}

	bytes := ip.To4()
	if bytes == nil {
		bytes = ip.To16()
	}

	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return bytes.String()
}
