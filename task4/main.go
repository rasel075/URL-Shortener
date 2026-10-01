package main

import (
  "fmt"
  "net/http"
  "sync"
  "time"
)

var (
  mu     sync.Mutex
  tokens = 5
  last   = time.Now()
)

func limit(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    mu.Lock()

    if time.Since(last) >= time.Second {
      tokens = 5
      last = time.Now()
    }

    if tokens <= 0 {
      mu.Unlock()
      http.Error(w, "Too Many Requests", 429)
      return
    }

    tokens--
    mu.Unlock()

    next.ServeHTTP(w, r)
  })
}

func main() {
  http.Handle("/", limit(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello!")
  })))

  http.ListenAndServe(":8080", nil)
}