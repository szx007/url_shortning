package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math"
	"strconv"
	"strings"
)

const (
	VAL   = 0x00000002
	INDEX = 0x0000003D
)

var base = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func Base62encode(num int) string {
	baseStr := ""
	for {
		if num <= 0 {
			break
		}

		i := num % 62
		baseStr += base[i]
		num = (num - i) / 62
	}
	return baseStr
}

func insert(slice *[]string, index int, value string) {
	rear := append([]string{}, (*slice)[index:]...)
	*slice = append(append((*slice)[:index], value), rear...)
}

func Transform(longURL string) string {
	md5Str := GetMd5String(longURL)

	tempSubStr := md5Str[0:8]
	n, _ := strconv.ParseInt(tempSubStr, 16, 64)
	var e int64
	v := make([]string, 0)

	for i := 0; i < 8; i++ {
		x := INDEX & n
		e |= ((VAL & n) >> 1) << uint(i)
		insert(&v, 0, base[x])
		n = n >> 6
	}
	e |= n << 5
	insert(&v, 0, base[e&INDEX])
	return strings.Join(v, "")

}

func Base62decode(base62 string) int {
	rs := 0
	len := len(base62)
	f := flip(base)
	for i := 0; i < len; i++ {
		rs += f[string(base62[i])] * int(math.Pow(62, float64(i)))
	}
	return rs
}

func flip(s []string) map[string]int {
	f := make(map[string]int)
	for index, value := range s {
		f[value] = index
	}
	return f
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
