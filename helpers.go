package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// FindSubIndex finds index
func FindSubIndex(user User, name string) (int, bool) {
	for i, v := range user.RSSSubs {
		if v.RSSname == name {
			return i, true
		}
	}
	return 0, false
}

// RemoveFromSubs list remover
func RemoveFromSubs(user *User, index int) {
	user.RSSSubs = append(user.RSSSubs[:index], user.RSSSubs[index+1:]...)
}

// FindRssIndex finds index
func FindRssIndex(user User, name string) (int, bool) {
	for i, v := range user.RSSList {
		log.Println(i, v.FetchName)
		if v.FetchName == name {
			return i, true
		}
	}
	return 0, false
}

// RemoveFromRss list remover
func RemoveFromRss(user *User, index int) {
	user.RSSList = append(user.RSSList[:index], user.RSSList[index+1:]...)
}

// CheckRSSFeed Rss Feed Checker
func CheckRSSFeed(user *User, fetch string) bool {
	for _, v := range user.RSSList {
		if v.FetchName == fetch {
			return true
		}
	}
	return false
}

// AddRSS Rss Address Add
func AddRSS(user *User, req JSONReq) {
	// add to subs
	newSub := RSSSub{
		RSSURL:  req.URL,
		RSSname: req.Name,
	}
	user.RSSSubs = append(user.RSSSubs, newSub)
}

// AddName Username add/change
func AddName(user *User, name string) {
	user.Username = name
}

// GetURLs Get All RSS data
func GetURLs(ptrUser *User) {
	var newRSSList []RSS
	ptrUser.RSSList = newRSSList
	for _, v := range ptrUser.RSSSubs {
		var newInfo RSSInfo
		var newRSS RSS
		raw := GetRss(v.RSSURL)
		UnpackRss(raw, &newInfo)
		RecieveRSS(newInfo, &newRSS, v.RSSname)
		ptrUser.RSSList = append(ptrUser.RSSList, newRSS)
	}
}

// GetURL Get One RSS data // ! refresh maybe
func GetURL(ptrUser *User, user User, name string) {
	// todo check if rss feed already exists

	for _, v := range ptrUser.RSSSubs {
		if v.RSSname == name {
			var newInfo RSSInfo
			var newRSS RSS

			index, ok := FindRssIndex(user, v.RSSname)

			if ok == true {
				raw := GetRss(v.RSSURL)
				UnpackRss(raw, &newInfo)
				RecieveRSS(newInfo, &newRSS, v.RSSname)
				ptrUser.RSSList[index] = newRSS
			} else {
				raw := GetRss(v.RSSURL)
				UnpackRss(raw, &newInfo)
				RecieveRSS(newInfo, &newRSS, v.RSSname)
				ptrUser.RSSList = append(ptrUser.RSSList, newRSS)
			}
		}
	}
}

// Response writer
func Response(success bool, msg string, w http.ResponseWriter) {
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
