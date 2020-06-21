package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

func main() {
	res, diagErr := netlink.SocketDiagTCPInfo(unix.AF_INET)
	if diagErr != nil {
		panic(diagErr)
	}

	debug := res[0]
	fmt.Printf("%+v\n", debug.InetDiagMsg)
	fmt.Printf("%+v\n", debug.TCPInfo)
}
