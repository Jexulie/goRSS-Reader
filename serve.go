package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// JSONReq request for rss reg
type JSONReq struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// JSONResp general response
type JSONResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

var user User

// Rss Feed Checker
func checkRSSFeed(user *User, fetch string) bool {
	for _, v := range user.RSSList {
		if v.FetchName == fetch {
			return true
		}
	}
	return false
}

// Rss Address Add
func addRSS(user *User, req JSONReq) {
	// add to subs
	newSub := RSSSub{
		RSSURL:  req.URL,
		RSSname: req.Name,
	}
	user.RSSSubs = append(user.RSSSubs, newSub)
}

// Username add/change
func addName(user *User, name string) {
	user.Username = name
}

// Get All RSS data
func getUrls(user *User) {
	// todo check if rss feed already exists
	for _, v := range user.RSSSubs {
		var newInfo RSSInfo
		var newRSS RSS
		raw := GetRss(v.RSSURL)
		UnpackRss(raw, &newInfo)
		RecieveRSS(newInfo, &newRSS, v.RSSname)
		user.RSSList = append(user.RSSList, newRSS)
	}
}

// Get One RSS data // ! refresh maybe
func getUrl(user *User, name string) {
	// ! set RSSList to empty array
	user.RSSList = nil
	// todo check if rss feed already exists
	for _, v := range user.RSSSubs {
		if v.RSSname == name {
			var newInfo RSSInfo
			var newRSS RSS
			raw := GetRss(v.RSSURL)
			UnpackRss(raw, &newInfo)
			RecieveRSS(newInfo, &newRSS, v.RSSname)
			user.RSSList = append(user.RSSList, newRSS)
		}
	}
}

// response writer
func response(success bool, msg string, w http.ResponseWriter) {
	resp := JSONResp{
		Success: success,
		Message: msg,
		Time:    time.Now().String(),
	}
	jsonize, _ := json.Marshal(resp)
	// ! General app/json headers
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Accept-Charset", "utf-8,Windows-1252")
	fmt.Fprintln(w, string(jsonize))
}

// IndexHandler index
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	response(true, "Welcome To Server", w)
	return
}

// AddUserHandler getter
func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//* addUser?name=john
		name, ok := r.URL.Query()["name"]
		if !ok || len(name) < 1 {
			response(false, "No Parameters Provided ...", w)
			return
		}
		addName(&user, name[0])

		response(true, "Username Added/Changed ...", w)
		// fmt.Printf("%+v \n", user)
		return
	}
	response(false, "Wrong End Point", w)
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
		addRSS(&user, req)

		response(true, "RSS Address Added ...", w)
		// log.Printf("%+v \n", user)
		return
	}
	response(false, "Wrong End Point", w)
	return
}

// GetterHandler getter
func GetterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		name, ok := r.URL.Query()["name"]
		if !ok || len(name) < 1 {
			response(false, "No Parameters Provided ...", w)
			return
		}
		getUrl(&user, name[0])
		// todo get rss if query is 0 get all
		res, _ := json.Marshal(&user)
		response(true, string(res), w)
		return
	}
	response(false, "Wrong End Point", w)
	return
}

// GetterAllHandler getter
func GetterAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getUrls(&user)
		// log.Printf("%+v \n", user)
		res, _ := json.Marshal(&user)
		response(true, string(res), w)
		// get rss if query is 0 get all
		return
	}
	response(false, "Wrong End Point", w)
	return
}

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
