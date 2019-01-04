package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var user User

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/addUser", AddUserHandler)
	r.HandleFunc("/addRSS", AdderHandler)
	r.HandleFunc("/getRSS", GetterHandler)
	r.HandleFunc("/getallRSS", GetterAllHandler)

	server := &http.Server{
		Addr:         "0.0.0.0:3333",
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 5,
		Handler:      r,
	}

	log.Println("Server Started ...")
	err := server.ListenAndServe()
	Check(err)
}

// * json.RawMessage -> new structs for json instead of having one for both
