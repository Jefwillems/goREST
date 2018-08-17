package main

import (
	"./db"
	"./models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Index(writer http.ResponseWriter, r *http.Request) {
	users := db.GetUsers(db.GetInstance())
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(users); err != nil {
		panic(err)
	}
}

func AddUser(at http.ResponseWriter, r *http.Request) {
	data := db.GetInstance()
	var m models.User
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &m)
	log.Print(m)
	db.AddUser(data, m)
	at.WriteHeader(http.StatusOK)
	at.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(at).Encode(m); err != nil {
		panic(err)
	}
}
