package main

import (
	"encoding/json"
	"net/http"
)

func Index(writer http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Name: "Jef Willems"},
		User{Name: "Frank"},
	}
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(users); err != nil {
		panic(err)
	}
}
