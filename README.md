# go-mproxy

Dead simple multi-host/round-robin reverse proxy.

## Usage

Just like this:

```go
package main

import (
  "log"
  "net/http"

  "github.com/atipugin/go-mproxy"
)

func main() {
  // set up backend hosts registry
  r, err := mproxy.NewRegistry([]string{
    "http://localhost:8081",
    "http://localhost:8082",
    "http://localhost:8083",
  })

  if err != nil {
    log.Fatalf("Error: %s", err)
  }

  // init proxy
  p := mproxy.NewReverseProxy(r)

  // cross fingers and start
  http.Handle("/", p)
  http.ListenAndServe(":8080", nil)
}
```

## TODO

- Add tests
- Add more load-balancing strategies
