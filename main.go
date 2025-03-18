package main

import (
	"adeptus-limitarius/framework"
	"log"
	"net/url"
)

func main() {
	targetURL := "http://localhost:3000/"
	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatal(err)
	}

	framework.StartLimiterServer(target)
}
