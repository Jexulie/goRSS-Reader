package main

import (
	"database/sql"
	"encoding/json"

	_ "github.com/lib/pq"
)

const connStr = "postgres://bossjs:bluewhale@192.168.0.112/bossdb?sslmode=disable"

// AddUserToDB user
func AddUserToDB(username string) {
	query := "INSERT INTO RSS_USER_LIST (USERNAME) VALUES ($1)"
	db, err := sql.Open("postgres", connStr)
	Check(err)
	defer db.Close()
	db.QueryRow(query, username)
}

// AddRssListToDB list
func AddRssListToDB(username string, list []RSSSub) {
	query := "UPDATE RSS_USER_LIST SET RSS_LIST = $1 WHERE USERNAME = $2"
	//* list (struct) -> (json)
	db, err := sql.Open("postgres", connStr)
	Check(err)
	defer db.Close()
	jsonize, _ := json.Marshal(&list)
	db.QueryRow(query, jsonize, username)
}

// GetUserListFromDB subs
func GetUserListFromDB(username string) RSSSub {
	query := "SELECT RSS_LIST FROM RSS_USER_LIST WHERE USERNAME = $1"
	db, err := sql.Open("postgres", connStr)
	Check(err)
	defer db.Close()
	var result RSSSub
	db.QueryRow(query, username).Scan(&result)
	return result
}
