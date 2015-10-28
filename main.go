package main

import (
	"fmt"
	"log"
)

var (
	counter = NewCounter(int64(100000))

	longUrls = []string{
		"https://golang.org",
		"https://www.baidu.com",
		"http://blog.jobbole.com/",
		"http://baidu.com",
		"http://www.vaikan.com/",
	}
)

func main() {
	fmt.Println("start...")
	counter.Run()

	i := len(longUrls)
	for i > 0 {
		i--
		longUrl := longUrls[i]
		su, err := ShortUrlCreate(longUrl)
		if err != nil {
			log.Fatalf("Short Url Create error, %v", err)
		}
		log.Printf("short url: %v", su)

        log.Printf("id: %d, su.Id: %d", GenId(su.Slug), su.Id)
	}


	fmt.Println("stop...")
}
