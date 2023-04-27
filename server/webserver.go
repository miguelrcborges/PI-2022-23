package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"database/sql"
)

func streamDevicesData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")

	f, ok := w.(http.Flusher)
	if ok {
		f.Flush()
	} else {
		fmt.Fprint(w, "Not able to receive stream.\n")
		return
	}

	for i := 0; true; i++ {
		select {
		case <-r.Context().Done():
			return
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
		fmt.Fprint(w, "Not able to receive stream.\n")
		return
	}

	for i := 0; true; i++ {
		select {
		case <-r.Context().Done():
			return
		default:
			fmt.Fprintf(w, "id: %d\nevent: updateDevicesCount\ndata: %d\n\n", i, len(devices))
			f.Flush()
			time.Sleep(5 * time.Second)
		}
	}
}

func queryUsers(w http.ResponseWriter, r *http.Request) {
	var query UsersQuery
	var rows *sql.Rows
	var err error

	params := r.URL.Query()
	if search := params.Get("s"); search != "" {
		wildcarded := "%" + search + "%";
		rows, err = db.Query("Select name, number from users where name like ? or number like ?;", wildcarded, wildcarded);
	} else {
		rows, err = db.Query("Select name, number from users;")
	}
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.Name, &user.Number)
		query.Users = append(query.Users, user)
	}

	// json , err := json.Marshal(query)
	err = json.NewEncoder(w).Encode(query)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	// fmt.Fprint(w, string(json))
}

func setUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	
	ip := params.Get("ip")

	if _, ok := devices[ip]; !ok {
		fmt.Fprintln(w, "Device not found")
		return
	}

	user_number := params.Get("un")

	row := db.QueryRow("select name, number from users where number = ?;", user_number)
	row.Scan(&devices[ip].UserName, &devices[ip].UserNumber)
	devices[ip].Order = "Waiting for assignment"
	devices[ip].orderReceived = 0;

	fmt.Fprint(w, "Success")
}

func setOrder(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	ip := params.Get("ip")
	
	if _, ok := devices[ip]; !ok {
		fmt.Fprintln(w, "Device not found")
		return
	}

	devices[ip].Order = params.Get("o")
	devices[ip].orderReceived = 0

	fmt.Fprint(w, "Success")
}
