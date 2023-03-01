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
	StatusCode int `json:"statusCode"`
}

// NOTE: ねこ画像取得API
const URL = "https://api.thecatapi.com/v1/images/search"

func main() {
  // ローカルで実行する場合
//   sub_main()

  // lambdaで実行する場合
  lambda.Start(hello)
//   parrot()
}

func sub_main() {
  url := get_url()
  fmt.Println(url)
  send_linebot(url)
}

func hello() (Response, error) {
  sub_main()
	return Response{StatusCode: 200}, nil
}

func parrot() {
  http.HandleFunc("/", lineHandler)
  log.Fatal(http.ListenAndServe(":8000", nil))
}

func lineHandler(w http.ResponseWriter, r *http.Request) {
  msg := "Hello World!!!!"
  fmt.Fprintf(w, msg)

  godotenv.Load(".env")
  bot, err := linebot.New(os.Getenv("LINE_BOT_CHANNEL_SECRET"),os.Getenv("LINE_BOT_CHANNEL_TOKEN"))
  if err != nil {
    log.Fatal(err)
  }

  // リクエストからBOTのイベントを取得
  events, err := bot.ParseRequest(r)
  if err != nil {
    if err == linebot.ErrInvalidSignature {
      w.WriteHeader(400)
    } else {
      w.WriteHeader(500)
    }
    return
  }
  for _, event := range events {
    // イベントがメッセージの受信だった場合
    if event.Type == linebot.EventTypeMessage {
      url := get_url()
      // NOTE: 画像メッセージ生成
      responseImage := linebot.NewImageMessage(url, url)

      // NOTE: 友達登録しているユーザー全員に画像を配信する
      if _, err := bot.BroadcastMessage(responseImage).Do(); err != nil {
        log.Fatal(err)
      }
    }
  }
  return
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
