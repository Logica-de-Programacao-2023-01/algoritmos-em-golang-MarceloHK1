package main

import "fmt"

func main() {
	var x, y float64
	fmt.Print("qual o valor de x? ")
	fmt.Scan(&x)
	fmt.Print("qual o valor de y? ")
	fmt.Scan(&y)

	if x > y {
		fmt.Print(x)
	} else {
		fmt.Println(y)
	}
}
