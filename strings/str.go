package main

import (
	"fmt"
)

func main() {
	str := "this is a string"
	fmt.Println(reverse(str))
	fmt.Println(Reverse(str))
}

func reverse(s string) string {
	runes := []rune(s)
	for l, r := 0, len(runes)-1; l <= r; l, r = l+1, r-1 { // this for loop does multiple assignment for iterators and post
		runes[l], runes[r] = runes[r], runes[l]
	}
	return string(runes)
}

func Reverse(s string) string {
	l, r := 0, len(s)-1
	xb := []byte(s)
	for {
		if l > r {
			return string(xb)
		}
		xb[l], xb[r] = xb[r], xb[l]
		l++
		r--
	}
}
