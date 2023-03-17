package main

import "fmt"

func main() {
	var x1, x2, x3 int
	fmt.Print("qual é x1? ")
	fmt.Scan(&x1)
	fmt.Print("qual é o x2? ")
	fmt.Scan(&x2)
	fmt.Print("qual é o x3? ")
	fmt.Scan(&x3)

	soma := x1 + x2 + x3

	fmt.Println("a soma é ", soma)
}
