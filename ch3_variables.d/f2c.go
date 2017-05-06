package main

import "fmt"

func main() {
    fmt.Println("Convert Fahrenheit to Celsius")
    fmt.Print("Enter temperature (F): ")
    var fahrenheit float64
    fmt.Scanf("%f", &fahrenheit)

    celsius := (fahrenheit - 32) * 5/9
    fmt.Println(celsius, "C")
}
