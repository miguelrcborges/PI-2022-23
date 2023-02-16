package main

type position struct {
	x, y float64
}

type userDetails struct {
	position position
	target position
	angle float64
}

type devicesMap map[int64]*userDetails
