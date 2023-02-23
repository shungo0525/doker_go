package controllers

import (
  "fmt"
  "net/http"
  "encoding/json"
  _ "github.com/go-sql-driver/mysql"

  "app/database"
  "app/models"
  "app/mail"
)

func HandleRequests() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/user", userPage)
  http.HandleFunc("/db_check", dbCheck)
  http.HandleFunc("/send_email", sendEmail)
  http.HandleFunc("/get_parameters", getParameters)
}

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Home")
  fmt.Println("Endpoint Hit: homePage")
}

func userPage(w http.ResponseWriter, r *http.Request){
  user := model.User{1, "user-1"}
  res, err := json.Marshal(user)
  if err != nil {
    panic(err)
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(res)

  fmt.Println("Endpoint Hit: userPage")
}

func getParameters(w http.ResponseWriter, r *http.Request) {
  if err := r.ParseForm(); err != nil {
    fmt.Println("parameter parse error")
  }

  for k, v := range r.Form {
    fmt.Println(k, v)
  }
  fmt.Fprintf(w, "Get Parameters")
}

func dbCheck(w http.ResponseWriter, r *http.Request){
  database.ConnectCheck()
  fmt.Fprintf(w, "DB check")
  fmt.Println("Endpoint Hit: db check")
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
  mail.SendEmail()
  fmt.Fprintf(w, "sen demail")
  fmt.Println("Endpoint Hit: send email")
}