package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("informe o valor de x e y ")
	fmt.Scan(&x, &y)

	if x > 0 && y > 0 {
		fmt.Print(x * y)
	} else {
		fmt.Print(x + y)
	}
}
