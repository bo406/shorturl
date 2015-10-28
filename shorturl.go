package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrCounter = errors.New("Counter error")
)

type ShortUrl struct {
	Id   int64
	Slug string
	URL  string
}

func (s *ShortUrl) String() string {
	return fmt.Sprintf("%s, %s", s.Slug, s.URL)
}

func ShortUrlCreate(url string) (*ShortUrl, error) {
	id := <-counter.ch
	if id <= 0 {
		return nil, ErrCounter
	}

	s := &ShortUrl{Id: id, URL: url}

	s.Slug = GenKey(s.Id)
	log.Printf("slug: %s", s.Slug)
	return s, nil
}
