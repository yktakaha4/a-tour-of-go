package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ipAddr IPAddr) String() string {
	octets := make([]string, 0)
	for _, a := range ipAddr {
		octets = append(octets, strconv.Itoa(int(a)))
	}
	return strings.Join(octets, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
