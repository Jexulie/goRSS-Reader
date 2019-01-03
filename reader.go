package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// GetRss gets Rss from web
func GetRss(url string) []byte {
	resp, err := http.Get(url)
	Check(err)

	data, err := ioutil.ReadAll(resp.Body)
	Check(err)
	return data
}

// UnpackRss to struct
func UnpackRss(data []byte, info *RSSInfo) {
	xml.Unmarshal(data, info)
	// ! problem starts after here
}

// RecieveRSS rss giver
func RecieveRSS(info RSSInfo, rss *RSS, fetch string) {
	rss.LastFetchTime = GetTimeNow()
	rss.FetchName = fetch
	rss.RSSInfo = info
}

// https://golang.org/pkg/encoding/json/#RawMessage

// func main() {
// 	news := "http://www.milliyet.com.tr/rss/rssNew/SonDakikaRss.xml"
// 	newsb := GetRss(news)
// 	var n RSS
// 	xml.Unmarshal(newsb, &n)
// 	fmt.Println(n)
// }
