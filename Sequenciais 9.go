package main

import "fmt"

func main() {
	var preço float64
	fmt.Print("qual o preço do produto? ")
	fmt.Scan(&preço)

	valorfinal := preço * 0.9

	fmt.Println("o preço final apos o desconto ser aplicado é ", valorfinal)
}
