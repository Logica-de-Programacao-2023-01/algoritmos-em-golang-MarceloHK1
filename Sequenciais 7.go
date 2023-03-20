package main

import "fmt"

func main() {
	var x int
	fmt.Print("qual valor de x? ")
	fmt.Scan(&x)

	antecessor := x - 1
	sucessor := x + 1

	fmt.Println("o antecessor e o sucessor de x são, respectivamente, ", antecessor, "e", sucessor)