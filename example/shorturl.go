package main

import (
	"github.com/bo406/shorturl"
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
	log.Println("start...")

	counter := shorturl.NewCounter(int64(100000))
	counter.Run()
    i := len(longUrls)
    for i > 0 {
        i--
        longUrl := longUrls[i]
        su, err := shorturl.ShortUrlCreate(longUrl, counter.GetCount())
        if err != nil {
            log.Fatalf("Short Url Create error, %v", err)
        }
        log.Printf("short url: %v", su)
        log.Printf("id: %d, su.Id: %d", shorturl.GenId(su.Slug), su.Id)
    }

	log.Println("stop...")
}
