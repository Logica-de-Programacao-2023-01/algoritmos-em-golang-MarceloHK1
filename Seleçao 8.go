package main

import "fmt"

func main() {
	var x int
	fmt.Print("qual o valor de x")
	fmt.Scan(&x)
	if x == 1 {
		fmt.Print("domingo")
	} else if x == 2 {
		fmt.Print("segunda")
	} else if x == 3 {
		fmt.Print("terça")
	} else if x == 4 {
		fmt.Print("quarta")
	} else if x == 5 {
		fmt.Print("quinta")
	} else if x == 6 {
		fmt.Print("sexta")
	} else {
		fmt.Print("sabado")
	}
}
