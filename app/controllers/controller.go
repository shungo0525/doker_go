package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  _ "github.com/go-sql-driver/mysql"

  "app/database"
  "app/api/gorm/models"
)

func main() {
  handleRequests()
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/user", userPage)
  http.HandleFunc("/db_check", dbCheck)

  err := http.ListenAndServe(":8000", nil)
  if err != nil {
    log.Fatal(err)
  }
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

func dbCheck(w http.ResponseWriter, r *http.Request){
  database.ConnectCheck()
  fmt.Fprintf(w, "DB check")
  fmt.Println("Endpoint Hit: db check")
}