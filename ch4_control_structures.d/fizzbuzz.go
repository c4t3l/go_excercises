package main

import "fmt"

func main() {
	/*
	   Print 1 to 100. For multiples of 3 print "fizz",
	   for multiples of 5 print "buzz". For numbers that
	   are both print "fizzbuzz"
	*/

	for i := 1; i <= 100; i++ {
		switch i % 3 {
		case 0:
			fmt.Println("fizz")
		case 1:
			fmt.Println(i)
		}
	}

}
