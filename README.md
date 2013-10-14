#Traffic Airbrake Middleware

Package airbrake implements a Traffic Middleware for [Airbrake](http://airbrake.io).

This is a Middleware for [Traffic](https://github.com/pilu/traffic).
It is base on [@tobi](https://github.com/tobi)'s [Airbrake library](https://github.com/tobi/airbrake-go).

## Example:

```go
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
```
