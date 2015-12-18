package shorturl

import (
	"errors"
	"fmt"
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
	return fmt.Sprintf("%d, %s, %s", s.Id, s.Slug, s.URL)
}

func ShortUrlCreate(url string, id int64) (*ShortUrl, error) {
	if id <= 0 {
		return nil, ErrCounter
	}

	s := &ShortUrl{Id: id, URL: url}
	s.Slug = GenKey(s.Id)

	return s, nil
}
