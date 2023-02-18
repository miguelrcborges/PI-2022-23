package main

import (
	"fmt"
	"net/http"
)

var devices devicesMap;

func main() {
	devices = make(devicesMap)

	http.HandleFunc("/api/get/data", streamDevicesData)
	http.HandleFunc("/api/get/quantity", streamAmountOfDevicesConnected)
	// http.HandleFunc("/api/new/device/", newDevice) /* /api/new/device/id/target_x/target_y */
	// http.HandleFunc("/api/delete/device/", deleteDevice) /* /api/delete/device/id */
	http.Handle("/", http.FileServer(http.Dir("static")))

	go startUdp()
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Print("Err: ", err.Error())
	}
}
