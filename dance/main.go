package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		f, err := os.Open("goph.txt")
		if err != nil {
			fmt.Println("error:", err)
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			time.Sleep(time.Millisecond * 100)
			fmt.Println(scanner.Text())
		}
		f.Close()
	}
}
