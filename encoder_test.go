package main

import "testing"

func TestEncoder(t *testing.T) {
	counter := NewCounter(int64(100000))
	counter.Run()
	n := <-counter.ch
	s := generateURL(n)
	if s != "F5E" {
		t.Fatalf("gen key error, id: %d, %s, %s", n, s, "F5E")
	}

	id := decodeURL(s)
	if id != 100000 {
		t.Fatalf("gen id error, %d, %d", id, 100000)
	}

}

func BenchmarkEncoder(b *testing.B) {
	b.StartTimer()
	counter := NewCounter(int64(100000))
	counter.Run()
	for i := 0; i < b.N; i++ {
		n := <-counter.ch
		s := generateURL(n)
		decodeURL(s)
	}
}
