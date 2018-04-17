// exercise from https://tour.golang.org/methods/18
package main

import "fmt"

type IPAddr [4]byte

func (ipa IPAddr) String() string {
  // tried to find the equivalent of Python's join but it's more
  // complicated for a byte slice
	return fmt.Sprintf("%v.%v.%v.%v", ipa[0], ipa[1], ipa[2], ipa[3])
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
