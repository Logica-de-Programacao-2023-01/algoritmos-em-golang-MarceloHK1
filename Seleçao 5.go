package main

import "fmt"

func main() {
	var x int
	fmt.Print("informe o valor de x")
	fmt.Scan(&x)

	if x%3 == 0 && x%5 == 0 {
		fmt.Print("x é multiplo de 3 e 5")
	} else {
		fmt.Print("x não é múltiplo de 3 e 5")
	}
}
