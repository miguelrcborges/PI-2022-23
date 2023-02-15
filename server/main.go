package main

import (
	"fmt"
	"net/http"
	"time"
)

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

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "helo")
}

func main() {
	http.HandleFunc("/api/get/positions", streamDevicesPosition)
	http.HandleFunc("/", hello)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Print("Err: ", err.Error())
	}
}
