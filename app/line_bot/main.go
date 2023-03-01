package main

import (
  "io"
  "os"
  "log"
  "fmt"
  "net/http"
  "encoding/json"

  "github.com/line/line-bot-sdk-go/v7/linebot"
  "github.com/joho/godotenv"
  "github.com/aws/aws-lambda-go/lambda"
)

type Item struct {
  Id string `json:"id"`
  Url   string `json:"url"`
}

type Response struct {
	Status string `json:"Status"`
}

// NOTE: ねこ画像取得API
const URL = "https://api.thecatapi.com/v1/images/search"

func main() {
  // ローカルで実行する場合
//   sub_main()

  // lambdaで実行する場合
  lambda.Start(hello)
}

func sub_main() {
  url := get_url()
  fmt.Println(url)
  send_linebot(url)
}

func hello() (Response, error) {
  sub_main()
	return Response{Status: fmt.Sprintf("ok")}, nil
}

func get_url() (url string) {
  var items []Item

  res, err := http.Get(URL)
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
  return items[0].Url
}

func send_linebot(url string) {
  godotenv.Load(".env")

  // NOTE: LineBot生成
  bot, err := linebot.New(os.Getenv("LINE_BOT_CHANNEL_SECRET"),os.Getenv("LINE_BOT_CHANNEL_TOKEN"))

  if err != nil {
    log.Fatal(err)
  }

  // NOTE: 画像メッセージ生成
  responseImage := linebot.NewImageMessage(url, url)

  // NOTE: 友達登録しているユーザー全員に画像を配信する
  if _, err := bot.BroadcastMessage(responseImage).Do(); err != nil {
    log.Fatal(err)
  }
}
