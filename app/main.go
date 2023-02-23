package main

import (
  "log"
  "net/http"

  "app/controllers"
  "app/models"
)

func main() {
  // マイグレーション
  model.Migrate()

  controllers.HandleRequests()

  err := http.ListenAndServe(":8000", nil)
  if err != nil {
    log.Fatal(err)
  }
}
