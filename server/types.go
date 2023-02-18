package main

type position struct {
	X, Y float64
}

type deviceDetails struct {
	Position position
	Target position
	Angle float64
}

type devicesMap map[string]*deviceDetails
