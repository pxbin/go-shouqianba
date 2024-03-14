package main

import (
	"context"
	"log"

	"github.com/pxbin/go-shouqianba"
)

func main() {
	config := &shouqianba.Config{}

	client := shouqianba.NewClient(config)

	// terminal activate
	res, _, err := client.Terminal.Activate(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.BizResponse.TerminalSN)
	log.Println(res.BizResponse.TerminalKey)
}
