package main

import "sync"

type DB struct {
	users Users
}

var instance *DB
var once sync.Once

func GetInstance() *DB {
	once.Do(func() {
		instance = &DB{
			users: Users{},
		}
	})
	return instance
}

func AddUser(db *DB, user User) User {
	users := append(db.users, user)
	return users[len(users)-1]
}

func GetUsers(db *DB) Users {
	return db.users
}
