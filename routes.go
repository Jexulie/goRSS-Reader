package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// IndexHandler index
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
	return
}

// AddUserHandler getter
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//* addUser?name=john
		name, ok := r.URL.Query()["name"]
		if !ok || len(name) < 1 {
			Response(false, "No Parameters Provided ...", w)
			return
		}
		//* Add RSS
		AddName(&user, name[0])
		//* Add to DB
		AddUserToDB(name[0])

		Response(true, "Username Added/Changed ...", w)
		// fmt.Printf("%+v \n", user)
		return
	}
	Response(false, "Wrong End Point", w)
	return
}

// AdderHandler adder
func AdderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bBody, err := ioutil.ReadAll(r.Body)
		Check(err)
		defer r.Body.Close()

		var req JSONReq
		err = json.Unmarshal(bBody, &req)
		Check(err)
		//* Add RSS
		AddRSS(&user, req)
		//* Add to DB
		AddRssListToDB(user.Username, user.RSSSubs)

		fmt.Printf("%+v \n", user.RSSSubs)

		Response(true, "RSS Address Added ...", w)
		// log.Printf("%+v \n", user)
		return
	}
	Response(false, "Wrong End Point", w)
	return
}

// GetterHandler getter - Get One Rss
func GetterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//! gets username and checks db
		name, ok := r.URL.Query()["name"]
		if !ok || len(name) < 1 {
			Response(false, "No Name Parameter Provided ...", w)
			return
		}

		username, ok := r.URL.Query()["username"]
		if !ok || len(username) < 1 {
			Response(false, "No Username Parameter Provided ...", w)
			return
		}
		user.Username = username[0]
		GetURL(&user, user, name[0])
		AddRssInfoToDB(username[0], user.RSSList)
		res, _ := json.Marshal(&user)
		Response(true, string(res), w)
		return
	}
	Response(false, "Wrong End Point", w)
	return
}

// GetterAllHandler getter - Get All Rss
func GetterAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//! gets username and checks db
		username, ok := r.URL.Query()["username"]
		if !ok || len(username) < 1 {
			Response(false, "No Username Parameter Provided ...", w)
			return
		}
		user.RSSSubs = GetRssSubsFromDB(username[0])
		user.RSSList = GetRssInfoFromDB(username[0])
		GetURLs(&user)
		AddRssInfoToDB(username[0], user.RSSList)
		// log.Printf("%+v \n", user)
		res, _ := json.Marshal(&user)
		Response(true, string(res), w)
		// get rss if query is 0 get all
		return
	}
	Response(false, "Wrong End Point", w)
	return
}

// TODO Add user system
