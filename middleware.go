package airbrake

import (
  "errors"
  "github.com/pilu/traffic"
  api "github.com/tobi/airbrake-go"
)

type AirbrakeMiddleware struct {}

func (middleware *AirbrakeMiddleware) ServeHTTP(w traffic.ResponseWriter, r *traffic.Request, next traffic.NextMiddlewareFunc) {
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
    nextMiddleware.ServeHTTP(w, r, next)
  }
}

func New(apiKey string) *AirbrakeMiddleware {
  api.ApiKey = apiKey

  return &AirbrakeMiddleware{}
}
