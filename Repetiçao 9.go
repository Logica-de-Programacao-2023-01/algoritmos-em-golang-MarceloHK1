package main

import "fmt"

func main() {
	soma := 0
	quantidade := 0

	for {
		var x int
		fmt.Print("digite o valor de x: ")
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		soma += x
		quantidade++
	}

	if quantidade == 0 {
		fmt.Print("nenhum numero digitado")
	} else {
		media := float64(soma) / float64(quantidade)
		fmt.Println("a media é: ", media)
	}
}
