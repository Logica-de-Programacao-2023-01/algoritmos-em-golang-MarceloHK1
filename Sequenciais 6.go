package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("qual é o seu salario? ")
	
	ns := salario * 1.15
	fmt.Println("o seu novo salario é ", ns)
}
