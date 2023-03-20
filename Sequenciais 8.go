package main

import "fmt"

func main() {
	var dias, valor float64
	fmt.Print("quantos dias voce trabalhou? ")
	fmt.Scan(&dias)
	fmt.Print("quanto voce recebe por dia? ")
	fmt.Scan(&valor)

	salario := dias * valor

	fmt.Println("seu salario é ", salario)
}