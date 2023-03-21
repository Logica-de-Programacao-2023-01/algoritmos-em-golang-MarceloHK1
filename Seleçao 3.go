package main

import "fmt"

func main() {
	var x int
	fmt.Print("informe x ")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Print("x é par")
	} else {
		fmt.Print("x é ímpar")
	}
}

