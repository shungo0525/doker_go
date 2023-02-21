package main

import (
  "log"
  "net/http"
  "app/controllers"
)

func main() {
  controllers.HandleRequests()

  err := http.ListenAndServe(":8000", nil)
  if err != nil {
    log.Fatal(err)
  }
}
