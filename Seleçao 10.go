package main

import "fmt"

func main() {
	var idade int
	fmt.Print("qual a sua idade? ")
	fmt.Scan(&idade)

	if idade < 9 {
		fmt.Print("voce é mirim")
	} else if idade > 9 && idade <= 13 {
		fmt.Print("voce é infantil")
	} else if idade > 13 && idade <= 17 {
		fmt.Print("voce é juvenil")
	} else if idade >= 18 {
		fmt.Print("voce é adulto")
	}
}
