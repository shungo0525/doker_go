package controllers

import (
  "fmt"
  "net/http"
  "encoding/json"
  "path"
  "strconv"
  _ "github.com/go-sql-driver/mysql"

  "app/models"
  "app/mail"
)

func HandleRequests() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/send_email", sendEmail)
  http.HandleFunc("/get_parameters", getParameters)
  http.HandleFunc("/users/", handleUsersRequest)
}

func handleUsersRequest(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  switch r.Method {
  case "GET":
    getUsers(w, r)
  case "POST":
    postUser(w, r)
  case "PUT":
    putUser(w, r)
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

  w.Write(res)
  fmt.Println("Endpoint Hit: get users")
}

func postUser(w http.ResponseWriter, r *http.Request) {
  name := r.FormValue("name")

  user, err := model.PostUser(name)
  res, err := json.Marshal(user)

  if err != nil {
    panic(err)
  }

  w.Write(res)
  fmt.Println("Endpoint Hit: post user")
}

func putUser(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(path.Base(r.URL.Path))
  if err != nil {
    fmt.Println(err)
    w.WriteHeader(400)
    return
  }

  name := r.FormValue("name")
  user, err := model.PutUser(id, name)

  if user.Id == 0 {
    w.Write([]byte("update failure"))
    return
  }

  res, err := json.Marshal(user)
  if err != nil {
    panic(err)
  }

  w.Write(res)
  fmt.Println("Endpoint Hit: put user")
}

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Home")
  fmt.Println("Endpoint Hit: homePage")
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