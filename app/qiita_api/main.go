package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Item struct {
  Title string `json:"title"`
  Url   string `json:"url"`
}

var url = "https://qiita.com/api/v2/items?page=1&per_page=20"

func main() {
	var items []Item

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Error: status code", res.StatusCode)
		return
	}

	body, _ := io.ReadAll(res.Body)

	if err := json.Unmarshal(body, &items); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", items)
}