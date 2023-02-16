package main

import (
	"fmt"
	"net/http"
	"time"
	"strings"
	"strconv"
)

var devices devicesMap;

func streamDevicesPosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")

	f, ok := w.(http.Flusher)
	if ok {
		f.Flush()
	} else {
		fmt.Fprintf(w, "Not able to receive stream.\n")
		return
	}
	
	for i := 0; true; i++ {
		select {
		case <- r.Context().Done():
			return;
		default:
			fmt.Fprintf(w, "id: %d\n", i)
			fmt.Fprintf(w, "event: update\n")
			fmt.Fprintf(w, "data: %d\n\n", i)
			f.Flush()
			time.Sleep(time.Second)
		}
	}
}


func streamAmountOfDevicesConnected(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")

	f, ok := w.(http.Flusher)
	if ok {
		f.Flush()
	} else {
		fmt.Fprintf(w, "Not able to receive stream.\n")
		return
	}
	
	for i := 0; true; i++ {
		select {
		case <- r.Context().Done():
			return;
		default:
			fmt.Fprintf(w, "id: %d\nevent: updateDevicesCount\ndata: %d\n\n", i, len(devices))
			f.Flush()
			time.Sleep(time.Second * 5)
		}
	}
}


func newDevice(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	if len(path) <= 6 {
		fmt.Fprintln(w, "Invalid parameters")
		return;
	}

	id, err := strconv.ParseInt(path[4], 10, 64)
	if err != nil {
		fmt.Fprintln(w, "Invalid device id")
		return;
	}

	devices[id] = &userDetails{}
	devices[id].target.x, _ = strconv.ParseFloat(path[5], 64)
	devices[id].target.y, _ = strconv.ParseFloat(path[6], 64)
	fmt.Fprintln(w, "Updated with success")
}


func deleteDevice(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	
	if len(path) < 4 {
		fmt.Fprintln(w, "No id was given")
	}

	id, err := strconv.ParseInt(path[4], 10, 64)
	if err != nil {
		fmt.Fprintln(w, "Invalid device id")
		return;
	}
	fmt.Fprintln(w, "Request done with success")

	delete(devices, id)
}


func main() {
	devices = make(devicesMap)
	http.HandleFunc("/api/get/positions", streamDevicesPosition)
	http.HandleFunc("/api/get/quantity", streamAmountOfDevicesConnected)
	http.HandleFunc("/api/new/device/", newDevice) /* /api/new/device/id/target_x/target_y */
	http.HandleFunc("/api/delete/device/", deleteDevice) /* /api/delete/device/id */
	http.Handle("/", http.FileServer(http.Dir("static")))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Print("Err: ", err.Error())
	}
}
