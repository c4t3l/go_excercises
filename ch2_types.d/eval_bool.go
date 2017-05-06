package main

import "fmt"

func main() {
	fmt.Println("Evaluating the following expression: (true && false) || (false && true) || !(false && false)")

	a := (true && false) || (false && true) || !(false && false)
	fmt.Println(a)
}
