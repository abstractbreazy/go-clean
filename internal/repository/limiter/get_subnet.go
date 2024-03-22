package limiter

import (
	"net"
	"strings"
)

// Returns the subnet string by given configurations and IP.
func (l *Limiter) GetSubnet(ip string) (string, error) {
	address := strings.Join([]string{ip, l.mask}, "/")
	_, subnet, err := net.ParseCIDR(address)
	if err != nil {
		return "", err
	}

	return subnet.String(), nil
}
