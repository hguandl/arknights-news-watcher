package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

var weiboTests = [4]string{
	"tests/weibo/01-mblog-with-article.json",
	"tests/weibo/02-mblog-with-video.json",
	"tests/weibo/03-mblog-with-text.json",
	"tests/weibo/04-mblog-with-tag-and-pic.json",
}
var weiboIdx = 0

var akAnnoTests = [5]string{
	"tests/akanno/00-init.json",
	"tests/akanno/01-new-gacha.json",
	"tests/akanno/02-activity-end.json",
	"tests/akanno/03-placehold.json",
	"tests/akanno/04-dev-news.json",
}
var akAnnoIdx = 0

func weiboHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile(weiboTests[weiboIdx])
	log.Printf("Deliverd %v\n", weiboTests[weiboIdx])
	weiboIdx = (weiboIdx + 1) % 4
	w.Write(data)
}

func akAnnoHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadFile(akAnnoTests[akAnnoIdx])
	log.Printf("Deliverd %v\n", akAnnoTests[akAnnoIdx])
	akAnnoIdx = (akAnnoIdx + 1) % 5
	w.Write(data)
}

func main() {
	listenAddr := ":8088"
	http.HandleFunc("/weibo", weiboHandler)
	http.HandleFunc("/akanno", akAnnoHandler)

	log.Printf("Listen at %v\n", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
