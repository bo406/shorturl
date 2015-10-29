package main

import (
	"log"
	"sync/atomic"
)

type Counter struct {
	next int64
	ch   chan int64
}

func NewCounter(next_id int64) *Counter {
	return &Counter{
		next: next_id,
		ch:   make(chan int64),
	}
}

func (c Counter) Run() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Counter run loop is terminating.  Detail: %v\n", r)
			}
		}()
		for {
			if c.ch != nil {
				c.ch <- atomic.LoadInt64(&c.next)
				atomic.AddInt64(&c.next, 1)
			} else {
				log.Print("Counter is closed.")
			}
		}
	}()
}

func (c Counter) Close() {
	if c.ch != nil {
		close(c.ch)
		c.ch = nil
	}
}
