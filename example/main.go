package main

import (
	"fmt"

	examplepb "github.com/lucasclerissepro/protoc-gen-temporal/example/gen/go"
)

func main() {
  client := examplepb.NewEmailWorkerClient()
  fmt.Println(client)
}
