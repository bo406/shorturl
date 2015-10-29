package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

const ALPHABET string = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const BASE float64 = 62

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

func md5ShortUrl1(url string) string {
	h := md5.New()
	io.WriteString(h, url)
	hx := h.Sum(nil)
	var result [4]string
	s := hex.EncodeToString(hx)
	length := len(s) / 8
	for i := 0; i < length; i++ {
		subhex := s[i*8 : i*8+8]
		bit, _ := strconv.ParseInt(subhex, 16, 64)
		var chHex = 0x3FFFFFFF & bit
		var outchars = ""
		for i := 0; i < 6; i++ {
			//base 62 3D base32 1F
			//比base值少2
			val := 0x0000001F & chHex
			outchars += ALPHABET[val : val+1]
			chHex = chHex >> 5
		}
		result[i] = outchars
	}
	log.Println(result)

	return result[0]
}

//idx range (0-3)
func md5ShortUrl2(url string, idx int) string {
	md5Hex := md5.Sum([]byte(url))
	md5Hash := hex.EncodeToString(md5Hex[:])
	shortStr := ""
	subStr := md5Hash[idx*8 : (idx+1)*8]
	x, _ := strconv.ParseInt(subStr, 16, 0)
	x = x & 0x3fffffff
	for k := 0; k < 6; k++ {
		index := 0x0000003d & x
		shortStr += ALPHABET[index : index+1]
		x = x >> 5
	}
	return shortStr
}
