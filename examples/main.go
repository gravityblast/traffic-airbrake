package main

import (
  "os"
  "fmt"
  "time"
  "net/http"
  "github.com/pilu/traffic"
  "github.com/pilu/traffic-airbrake"
)

func rootHandler(w traffic.ResponseWriter, r *http.Request) {
  err := fmt.Sprintf("Error at %v", time.Now())
  panic(err)
}

func main() {
  traffic.SetVar("env", "production")
  router := traffic.New()
  router.AddMiddleware(airbrake.New(os.Getenv("AIRBRAKE_API_KEY")))

  // Routes
  router.Get("/", rootHandler)
  router.Run()
}
