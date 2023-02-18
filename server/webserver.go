package main

import (
	"fmt"
	"net/http"
	"time"
	"strings"
	"strconv"
	"encoding/json"
)

func streamDevicesData(w http.ResponseWriter, r *http.Request) {
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
			json, _ := json.Marshal(devices)
			fmt.Fprintf(w, "id: %d\nevent: updateDevicesData\ndata: %s\n\n", i, json)
			f.Flush()
			time.Sleep(5 * time.Second)
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
			time.Sleep(5 * time.Second)
		}
	}
}


func newDevice(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	if len(path) <= 6 {
		fmt.Fprintln(w, "Invalid parameters")
		return;
	}

	id := path[4]
	devices[id] = &deviceDetails{}
	devices[id].Target.X, _ = strconv.ParseFloat(path[5], 64)
	devices[id].Target.Y, _ = strconv.ParseFloat(path[6], 64)
	fmt.Fprintln(w, "Updated with success")
}


func deleteDevice(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	
	if len(path) < 4 {
		fmt.Fprintln(w, "No id was given")
	}

	fmt.Fprintln(w, "Request done with success")

	delete(devices, path[4])
}
