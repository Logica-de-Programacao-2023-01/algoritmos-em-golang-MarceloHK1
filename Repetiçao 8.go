package main

import "fmt"

func main() {
	var x int
	fmt.Print("informe o valor de x ")
	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		if x%i != 0 {
			continue
		}
		fmt.Print(i, ",")
	}
}
