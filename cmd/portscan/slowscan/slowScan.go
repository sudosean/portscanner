package portscan

import (
	"fmt"
	"net"
)

func Slowscan() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Sprintf("Port %d may be closed or filtered", i)
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
