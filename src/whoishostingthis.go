package main 

import (
	"fmt"
	"net"
)

func main() {
	
	var hostToLookUp string = "kloody.com"
	
	fmt.Println("Host to look up: ", hostToLookUp)
	
	hostIPs, err := net.LookupIP(hostToLookUp)
	if err != nil {
		fmt.Println("The host ", hostToLookUp, " is not reachable")
	}
	fmt.Println("IPs: ", hostIPs)
	
	
}

