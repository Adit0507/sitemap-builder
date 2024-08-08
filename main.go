package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	urlFlag:= flag.String("url", "https://letterboxd.com/", "the url I want to build a sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)
	resp, err := http.Get(*urlFlag)
	if err!= nil{
		panic(err)
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

}