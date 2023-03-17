package main

import "fmt"

func main() {
	var x int
	fmt.Print("qual é o valor de x?")
	fmt.Scan(&x)

	dobro := 2 * x
	triplo := 3 * x
	quadruplo := 4 * x

	fmt.Println("o dobro é", dobro)
	fmt.Println("o triplo é", triplo)
	fmt.Println("o quadruplo é", quadruplo)
}
