package main

import "fmt"

func main() {
	var peso, altura float64
	fmt.Print("qual é o seu peso? ")
	fmt.Scan(&peso)
	fmt.Print("qual é a sua altura? ")
	fmt.Scan(&altura)

	imc := peso / (altura * altura)

	fmt.Println("seu imc é ", imc)
}
