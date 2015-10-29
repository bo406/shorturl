package main

import "testing"

func TestEncoder(t *testing.T) {
	//	counter := NewCounter(int64(100000))
	//	counter.Run()
	//	n := <-counter.ch
	n := int64(100000)
	s := generateURL(n)
	if s != "0aU" {
		t.Fatalf("gen key error, id: %d, %s, %s", n, s, "0aU")
	}

	id := decodeURL(s)
	if id != 100000 {
		t.Fatalf("gen id error, %d, %d", id, 100000)
	}

}

func BenchmarkEncoder(b *testing.B) {
	b.StartTimer()
	//	counter := NewCounter(int64(100000))
	//	counter.Run()
	n := int64(100000)
	for i := 0; i < b.N; i++ {
		s := generateURL(n)
		n++
		decodeURL(s)
	}
}
