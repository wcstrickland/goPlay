package main

import (
	"fmt"
)

func main() {
	fmt.Println(Nfibs(6))
}

func Nfibs(n int) int {
	f, s := 0, 1
	for i := 0; i < n-2; i++ {
		f, s = s, f+s
	}
	if n <= 1 {
		return 1
	} else {
		return s
	}
}
