package main

import (
	"net"
	"os"
	"fmt"
	//"strings"
)

func startUdp() {
	buf := make([]byte, 1024)
	udpserver, err := net.ListenPacket("udp", ":1050")
	if err != nil {
		fmt.Println("Error creating an UDP server at port 1050")
		os.Exit(2)
	}
	defer udpserver.Close()

	for {
		n, addr, err := udpserver.ReadFrom(buf)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}

		ip := strings.Split(addr.String(), ":")[0]

		if _, ok := devices[ip]; !ok {
			devices[ip] = &deviceDetails{}
		}

		fmt.Printf("%s: %s\n", ip, buf[:n])
	}
}
