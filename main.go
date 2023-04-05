package main

import (
	"context"
	"fmt"
	"github.com/trajan0x/moe-18demo/moe"
)

func main() {
	m := moe.New()
	host, err := m.Start(context.Background(), "http://localhost:8080")
	fmt.Println(host)
	if err != nil {
		panic(err)
	}
}
