package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		j := i % 3
		if j == 0 {
			fmt.Println(i)
		}
	}
}
