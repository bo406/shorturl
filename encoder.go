package main

/*
实现算法与key中的一样，效率稍低一些
BenchmarkEncoder-4	 2000000	       898 ns/op
BenchmarkGenKen-4 	 3000000	       555 ns/op
*/
import (
	"math"
	"strings"
)

const ALPHABET string = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const BASE float64 = 49

func generateURL(n int64) string {
	s := ""
	var num float64 = float64(n)

	for num > 0 {
		s = string(ALPHABET[int(num)%int(BASE)]) + s
		num = math.Floor(num / BASE)
	}

	return s
}

func decodeURL(s string) int {
	var num int = 0

	for _, element := range strings.Split(s, "") {
		num = num*int(BASE) + strings.Index(ALPHABET, element)
	}

	return num
}
