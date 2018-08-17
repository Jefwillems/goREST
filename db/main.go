package db

import "sync"
import "../models"

type DB struct {
	users models.Users
}

var instance *DB
var once sync.Once

func GetInstance() *DB {
	once.Do(func() {
		instance = &DB{
			users: models.Users{},
		}
	})
	return instance
}

func AddUser(db *DB, user models.User) models.User {
	users := append(db.users, user)
	return users[len(users)-1]
}

func GetUsers(db *DB) models.Users {
	return db.users
}
