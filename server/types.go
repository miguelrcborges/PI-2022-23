package main

import "time"

type deviceDetails struct {
	UserNumber  int64
	UserName    string
	Order       string
	lastRequest time.Time
}

type devicesMap map[string]*deviceDetails

type UsersQuery struct {
	Users []User;
}

type User struct {
	Name   string
	Number int64
}
