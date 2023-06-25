package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"strconv"
	"time"
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
			devices[ip] = &deviceDetails{
				UserName: "To assign",
				Order:    "Waiting for assignment",
			}
		}

		devices[ip].lastRequest = time.Now()
		go handleRequest(udpserver, ip, addr, buf[:n])
	}
}

func handleRequest(conn net.PacketConn, ip string, addr net.Addr, message []byte) {
	fmt.Println(ip, string(message))

	reply_status, _ := strconv.Atoi(string(message))
	if reply_status == 2 {
		devices[ip].orderReceived = 0
	}

	if devices[ip].orderReceived == 1 {
		return
	}

	devices[ip].orderReceived = int8(reply_status)

	if devices[ip].orderReceived == 1 {
		return
	}

	var buf []byte
	if devices[ip].UserName != "To assign" {
		names := strings.Split(devices[ip].UserName, " ")
		buf = []byte(names[0] + " " + names[len(names) - 1] + "\n" + devices[ip].Order)
	} else {
		buf = []byte(ip + "\n" + devices[ip].Order)
	}

	conn.WriteTo(buf, addr)
	fmt.Println("Sent to", ip, string(buf))
}
