package main

import (
	"fmt"
	"net/http"

	"database/sql"
	_ "modernc.org/sqlite"
)

var db *sql.DB
var devices devicesMap

func main() {
	var err error
	db, err = sql.Open("sqlite", "utentes.db")
	_ = db
	if err != nil {
		fmt.Println(err.Error())
	}

	devices = make(devicesMap)

	http.HandleFunc("/api/get/users", queryUsers)
	http.HandleFunc("/api/get/data", streamDevicesData)
	http.HandleFunc("/api/get/quantity", streamAmountOfDevicesConnected)
	http.Handle("/", http.FileServer(http.Dir("static")))

	go startUdp()
	go cleanup()
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Print("Err: ", err.Error())
	}
}
