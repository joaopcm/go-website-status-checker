package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string {
		"http://goole.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<- c)
	}
}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)

	if error != nil {
		c <- link + " might be down"
		return
	}

	c <- link + " is up!"
}
