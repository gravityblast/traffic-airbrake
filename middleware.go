package airbrake

import (
  "errors"
  "github.com/pilu/traffic"
  api "github.com/tobi/airbrake-go"
)

type AirbrakeMiddleware struct {}

func (middleware *AirbrakeMiddleware) ServeHTTP(w traffic.ResponseWriter, r *traffic.Request, next traffic.NextMiddlewareFunc) (traffic.ResponseWriter, *traffic.Request) {
  defer func() {
    if rec := recover(); rec != nil {
      if err, ok := rec.(error); ok {
        api.Error(err, r.Request)
      } else if err, ok := rec.(string); ok {
        api.Error(errors.New(err), r.Request)
      }
      panic(rec)
    }
  }()

  if nextMiddleware := next(); nextMiddleware != nil {
    w, r = nextMiddleware.ServeHTTP(w, r, next)
  }

  return w, r
}

func New(apiKey string) *AirbrakeMiddleware {
  api.ApiKey = apiKey

  return &AirbrakeMiddleware{}
}
