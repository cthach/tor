package tor

import (
	"context"
	"fmt"
	"net"
	"strings"
)

const dnsExitListHost = "ip-port.exitlist.torproject.org"

var answerExitNet = &net.IPNet{IP: net.IP{127, 0, 0, 0}, Mask: net.CIDRMask(24, 32)}

// IsExit returns nil if the IP address belongs to an exit node.
func IsExit(ctx context.Context, ip string) (bool, error) {
	if !isIP(ip) {
		return false, fmt.Errorf("%s is not a valid IP address", ip)
	}

	answers, err := queryDNSExitList(ctx, ip, "1.2.3.4", 80)
	if err != nil {
		if strings.Contains(err.Error(), "no such host") {
			return false, nil
		}
		return false, fmt.Errorf("error while querying Tor Exit List service: %v", err)
	}
	if len(answers) == 0 {
		return false, fmt.Errorf("no answers received")
	}
	if !isIP(answers[0]) {
		return false, fmt.Errorf("received an invalid IP for answers")
	}

	answerIP := net.ParseIP(answers[0])

	if !answerExitNet.Contains(answerIP) {
		return false, nil
	}

	return true, nil
}

// queryDNSExitList queries the Tor Project's DNS based node exit list and returns the response, if any.
var queryDNSExitList = func(_ context.Context, src string, dst string, port uint16) ([]string, error) {
	return net.LookupHost(
		fmt.Sprintf(
			"%s.%d.%s.%s",
			reverse(src),
			port,
			reverse(dst),
			dnsExitListHost,
		))
}
