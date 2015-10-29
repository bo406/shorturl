package shorturl

import "testing"

func TestGen(t *testing.T) {
	n := int64(100000)
	s := GenKey(n)
	if s != "0aU" {
		t.Fatalf("gen key error, n: %d, %s, %s", n, s, "0aU")
	}
	id := GenId(s)
	if id != n {
		t.Fatalf("gen id error, %d, %d", id, n)
	}
}

func BenchmarkGen(b *testing.B) {
	b.StartTimer()
	n := int64(100000)
	for i := 0; i < b.N; i++ {
		s := GenKey(n)
		GenId(s)
		n++
	}
}
