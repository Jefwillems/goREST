package main

import "time"

type User struct {
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
}

type Users []User
