package main

import "fmt"

func main() {
	var peso float64
	fmt.Print("qual seu peso em kg?")
	fmt.Scan(&peso)

	weight := peso * 2.205

	fmt.Println("seu peso em libras é ", weight)
}