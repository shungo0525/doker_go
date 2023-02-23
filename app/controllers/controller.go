package controllers

import (
  "fmt"
  "net/http"
  "encoding/json"
  _ "github.com/go-sql-driver/mysql"

  "app/models"
  "app/mail"
)

func HandleRequests() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/user", userPage)
  http.HandleFunc("/send_email", sendEmail)
  http.HandleFunc("/get_parameters", getParameters)
  http.HandleFunc("/users", handleUsersRequest)
}

func handleUsersRequest(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case "GET":
    getUsers(w, r)
  case "POST":
  case "PUT":
  case "DELETE":
  default:
    w.WriteHeader(405)
  }
}

func getUsers(w http.ResponseWriter, r *http.Request) {
  users, _ := model.GetUsers()

  res, err := json.Marshal(users)
  if err != nil {
    panic(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(res)

  fmt.Println("Endpoint Hit: get users")
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

func sendEmail(w http.ResponseWriter, r *http.Request) {
  mail.SendEmail()
  fmt.Fprintf(w, "sen demail")
  fmt.Println("Endpoint Hit: send email")
}