package main

var (
	keyChars  = []byte("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	decodeMap = makeDecodeMap()
)

func makeDecodeMap() map[byte]int64 {
	m := make(map[byte]int64)

	for i, b := range keyChars {
		m[b] = int64(i)
	}

	return m
}

func GenKey(n int64) string {
	if n == 0 {
		return string(keyChars[0])
	}

	l := int64(len(keyChars))
	s := make([]byte, 20)
	i := int64(len(s))

	for n > 0 && i >= 0 {
		i--
		j := n % l
		n = (n - j) / l
		s[i] = keyChars[j]
	}

	return string(s[i:])
}

func GenId(key string) int64 {
	l := int64(len(keyChars))
	n := int64(0)
	for _, b := range key {
		n *= l
		n += decodeMap[byte(b)]
	}

	return n
}
