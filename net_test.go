package utils


import (
	"fmt"
	"testing"
)


func TestHostname2IPv4(t *testing.T) {
	str := "127.0.0.1:100102"

	ip := Hostname2IPv4(str)

	fmt.Printf("ip string: %s\n", ip)

	str = "localhost:100102"
	ip = Hostname2IPv4(str)
	fmt.Printf("ip2 string: %s\n", ip)

	str = "localhost"
	ip = Hostname2IPv4(str)
	fmt.Printf("ip3 string: %s\n", ip)


	str = "127.0.0.1"
	ip = Hostname2IPv4(str)
	fmt.Printf("ip4 string: %s\n", ip)

	str = ":10102"
	ip = Hostname2IPv4(str)
	fmt.Printf("ip5 string: %s\n", ip)

}
