package main

import (
	"github.com/philipbo/shorturl"
	"log"
)

var (
	longUrls = []string{
		"https://golang.org",
		"https://www.baidu.com",
		"http://blog.jobbole.com/",
		"http://baidu.com",
		"http://www.vaikan.com/",
	}
)

func main() {
	counter := shorturl.NewCounter(int64(100000))
	counter.Run()
	i := len(longUrls)
	for i > 0 {
		i--
		longUrl := longUrls[i]
		su, err := shorturl.ShortUrlCreate(longUrl, counter.GetCount())
		if err != nil {
			log.Fatalf("ShortURL Create error, %v", err)
		}
		log.Printf("ShortURL: %v", su.String())
	}
}
