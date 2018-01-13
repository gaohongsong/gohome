package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"time"
)

type PageSize struct {
	url  string
	size int
}

func main() {
	/*
	[2018-01-13T13:58:33.605967+08:00]: http://www.google.com
	[2018-01-13T13:58:33.605967+08:00]: http://www.apple.com
	[2018-01-13T13:58:33.605967+08:00]: http://www.apple.com
	[2018-01-13T13:58:33.605967+08:00]: http://www.baidu.com
	{http://www.baidu.com 112326}
	{http://www.apple.com 45865}
	{http://www.apple.com 45865}
	{http://www.google.com 2477}
	*/
	urls := []string{
		"http://www.apple.com",
		"http://www.apple.com",
		"http://www.baidu.com",
		"http://www.google.com",
	}

	results := make(chan PageSize)
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("[%v]: %v\n", time.Now().Format(time.RFC3339Nano), url)
			if res, err := http.Get(url); err == nil {
				defer res.Body.Close()
				bs, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Fatal(err)
				}
				results <- PageSize{
					url:  url,
					size: len(bs),
				}
			} else {
				log.Fatal(err)

			}
		}(url)
	}
	for range urls {
		result := <-results
		fmt.Println(result)
	}
}
