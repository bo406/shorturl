package main

import "testing"

func TestGen(t *testing.T) {
	counter := NewCounter(int64(100000))
	counter.Run()
	n := <-counter.ch
	s := GenKey(n)
	if s != "0aU" {
		t.Fatalf("gen key error, n: %d, %s, %s", n, s, "0aU")
	}
	id := GenId(s)
	if id != 100000 {
		t.Fatalf("gen id error, %d, %d", id, 100000)
	}
}

func BenchmarkGen(b *testing.B) {
	b.StartTimer()
	counter := NewCounter(int64(100000))
	counter.Run()
	for i := 0; i < b.N; i++ {
		n := <-counter.ch
		s := GenKey(n)
		GenId(s)
	}
}
