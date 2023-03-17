package main

import "fmt"

func main() {
	var idade int
	fmt.Print("qual é a sua idade? ")
	fmt.Scan(&idade)

	id := idade * 365

	fmt.Println("a sua idade em dias é ", id)
}