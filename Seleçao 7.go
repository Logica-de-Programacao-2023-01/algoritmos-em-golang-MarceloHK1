package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("qual seu salario? ")
	fmt.Scan(&salario)

	if salario < 1000 {
		fmt.Print("seu novo salario é", 1.1*salario)
	} else {
		fmt.Print("seu novo salario é", 1.05*salario)
	}
}
