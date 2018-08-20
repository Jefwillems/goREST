package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Index(writer http.ResponseWriter, r *http.Request) {
	users := GetUsers(GetInstance())
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(users); err != nil {
		panic(err)
	}
}

func AddUserHandler(at http.ResponseWriter, r *http.Request) {
	data := GetInstance()
	var m User
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &m)
	log.Print(m)
	AddUser(data, m)
	at.WriteHeader(http.StatusOK)
	at.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(at).Encode(m); err != nil {
		panic(err)
	}
}
