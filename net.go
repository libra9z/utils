package utils

import (
	"strings"
	"net"
)

func Hostname2IPv4( hostn string ) ( ip string) {
	if hostn == "" {
		return ip
	}

	ss := strings.Split(hostn,":")
	if ss[0] == "" {

		ss[0] = "0.0.0.0"
	}
	s1,_ := net.LookupHost(ss[0])

	var p string

	for _,v := range s1 {
		if v != "::1" {
			p = v
			break
		}
	}

	if len(ss)>1 {
		ip = p +":"+ss[1]
	}else{
		ip = p
	}

	return
}