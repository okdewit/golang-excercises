package stringers

import "fmt"

type IPAddr [4]byte

var Hosts = map[string]IPAddr{
	"loopback":  {127, 0, 0, 1},
	"googleDNS": {8, 8, 8, 8},
}

func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}

