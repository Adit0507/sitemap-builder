package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag:= flag.String("url", "https://letterboxd.com/", "the url I want to build a sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)
}