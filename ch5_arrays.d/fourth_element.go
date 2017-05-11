package main

import "fmt"

func main() {
	fmt.Println("How do you access the fourth element of an array or slice?")
	fmt.Println("array = [ 00 45 66 89 23 99 90 ]")

	array := [7]int{
		00,
		45,
		66,
		89,
		23,
		99,
		90,
	}

	fmt.Println("Fourth element is", array[3])
}
