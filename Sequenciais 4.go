package main

import "fmt"

func main() {
	var x1, x2, x3 float64
	fmt.Print("qual o valor de x1? ")
	fmt.Scan(&x1)
	fmt.Print("qual o valor de x2? ")
	fmt.Scan(&x2)
	fmt.Print("qual o valor de x3? ")
	fmt.Scan(&x3)

	media := ((2 * x1) + (3 * x2) + (5 * x3)) / 10

	fmt.Println("a media é ", media)
}
