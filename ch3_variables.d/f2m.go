package main

import "fmt"

func main() {
    fmt.Println("Convert Feet to Meters")
    fmt.Print("Enter length (Ft): ")
    var feet float64
    fmt.Scanf("%f", &feet)

    meters := feet * 0.3048
    fmt.Println(meters, "M")
}
