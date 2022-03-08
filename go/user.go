package main

import "time"

type User struct {
	Id int `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	Lastname string `json:"lastName,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
	Email string `json:"email"`
	Password string `json:"password"`
}