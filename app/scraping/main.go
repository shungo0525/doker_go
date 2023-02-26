package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const URL = "https://shungoblog.com/"
const CLASS_NAME = ".entry-card-wrap"

func main() {
  url := get_url()
  fmt.Println(url)
}

func get_url() (url string) {
  res, err := http.Get(URL)
  if err != nil {
    log.Println("htmlの取得に失敗しました", err)
    return
  }
  defer res.Body.Close()

  // NOTE: スクレイピングしたURLを格納するスライス
  urls := make([]string, 0)

  doc, _ := goquery.NewDocumentFromReader(res.Body)
  doc.Find(CLASS_NAME).Each(func(i int, s *goquery.Selection) {
    url, _ := s.Attr("href")
    urls = append(urls, url)
  })

  rand.Seed(time.Now().UnixNano())
  index := rand.Intn(len(urls)-1)

  return urls[index]
}