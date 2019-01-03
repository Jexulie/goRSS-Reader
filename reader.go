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
func UnpackRss(data []byte) *RSSInfo {
	var newRSSInfo RSSInfo
	xml.Unmarshal(data, newRSSInfo)
	return &newRSSInfo
}

// RecieveRSS rss giver
func RecieveRSS(info *RSSInfo, url string) *RSS {
	var newRSS RSS
	newRSS.LastFetchTime = GetTimeNow()
	newRSS.FetchURL = url
	newRSS.RSSInfo = info
	return &newRSS
}

// func main() {
// 	news := "http://www.milliyet.com.tr/rss/rssNew/SonDakikaRss.xml"
// 	newsb := GetRss(news)
// 	var n RSS
// 	xml.Unmarshal(newsb, &n)
// 	fmt.Println(n)
// }
