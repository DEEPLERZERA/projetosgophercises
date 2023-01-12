package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "url to send request to")
	flag.Parse()

	fmt.Println(*urlFlag)
}
