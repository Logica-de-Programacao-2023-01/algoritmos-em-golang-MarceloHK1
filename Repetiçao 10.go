package main

import "fmt"

func main() {
	var x, maior int

	for {
		fmt.Print("digite o valor de x: ")
		fmt.Scan(&x)
		if x == 0 {
			break
		}

		if x > maior {
			maior = x
		}
	}

	fmt.Println("o maior numero é ", maior)
}
