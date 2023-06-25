package main

import "time"

func cleanup() {
	for {
		now := time.Now()
		for ip, data := range devices {
			if now.Sub(data.lastRequest).Minutes() > 1 {
				delete(un2ip, devices[ip].UserNumber)
				delete(devices, ip)
			}
		}
		time.Sleep(15 * time.Second)
	}
}
