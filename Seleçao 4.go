package main

import "fmt"

func main() {
	var peso, altura float64
	fmt.Print("qual seu peso e altura? ")
	fmt.Scan(&peso, &altura)

	imc := peso / (altura * altura)

	if imc < 18.5 {
		fmt.Print("voce esta abaixo do peso ideal")
	} else if imc > 30 {
		fmt.Print("voce esta acima do peso ideal")
	} else {
		fmt.Print("voce esta na faixa de peso ideal")
	}

}
