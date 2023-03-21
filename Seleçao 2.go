package main

import "fmt"

func main() {
	var x, y, z float64
	fmt.Print("quais os valores de x, y e z? ")
	fmt.Scan(&x, &y, &z)

	if x < y && x < z {
		fmt.Print(x)
	} else if y < x && y < z {
		fmt.Print(y)
	} else {
		fmt.Print(z)
	}
}
