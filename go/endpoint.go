package main

type Endpoint string

const (
	//Base Endpoints
	root  Endpoint = "/"
	login Endpoint = "/login"

	//User Endpoints
	home       Endpoint = "/user/home"
	votes      Endpoint = "/user/votes"
	ballots    Endpoint = "/user/ballots"
	attendance Endpoint = "/user/attendace"
	notes      Endpoint = "/user/notes"
	timetable  Endpoint = "/user/timetable"
	calendar   Endpoint = "/user/calendar"

	//API Endpoints
	api Endpoint = "/api"
)
