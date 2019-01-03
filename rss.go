package main

import (
	"log"
	"time"
)

// User server
type User struct {
	Username string   `json:"username"`
	RSSSubs  []RSSSub `json:"rss_subs"`
	RSS      []*RSS   `json:"rss_info_list"`
}

// RSSSub list
type RSSSub struct {
	RSSURL  string `json:"rss_url"`
	RSSname string `json:"rss"`
}

// RSS struct
type RSS struct {
	LastFetchTime string   `json:"last_fetch_time"`
	FetchURL      string   `json:"fetch_url"`
	RSSInfo       *RSSInfo `json:"rss_info"`
}

// RSSInfo xml
type RSSInfo struct {
	Channel Channel `xml:"channel" json:"channel"`
}

// Channel xml
type Channel struct {
	Title       string `xml:"title" json:"title"`
	Link        string `xml:"link" json:"link"`
	Description string `xml:"description" json:"description"`
	// Ops ...
	Language string `xml:"language" json:"language"`
	Docs     string `xml:"docs" json:"docs"`
	PubDate  string `xml:"pubDate" json:"pubDate"`
	GUID     string `xml:"guid" json:"guid"`
	Category string `xml:"category" json:"category"`
	Image    Image  `xml:"image" json:"image"`
	// etc...
	Item []Item `xml:"item" json:"item"`
}

// Item xml
type Item struct {
	Title       string `xml:"title" json:"title"`
	Link        string `xml:"link" json:"link"`
	Description string `xml:"description" json:"description"`
}

// Image xml
type Image struct {
	URL  string `xml:"url" json:"url"`
	Link string `xml:"link" json:"link"`
}

// Check Error Checker
func Check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// GetTimeNow time format
func GetTimeNow() string {
	return time.Now().String()
}
