package main

import (
	"context"
	"fmt"

	client "github.com/weaviate/weaviate-go-client/v4/weaviate"
)

func main() {
  config := client.Config{
    Scheme: "http",
    Host:   "localhost:8080",
  }
  c, err := client.NewClient(config)
  if err != nil {
    fmt.Printf("Error occurred %v", err)
    return
  }
  metaGetter := c.Misc().MetaGetter()
  meta, err := metaGetter.Do(context.Background())
  if err != nil {
    fmt.Printf("Error occurred %v", err)
    return
  }
  fmt.Printf("Weaviate meta information\n")
  fmt.Printf("hostname: %s version: %s\n", meta.Hostname, meta.Version)
  fmt.Printf("enabled modules: %+v\n", meta.Modules)
}