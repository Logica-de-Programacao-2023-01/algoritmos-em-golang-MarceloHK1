package main

import "fmt"

func main() {
	var x, y, z float64
	fmt.Print("informe os valores de x, y, z ")
	fmt.Scan(&x, &y, &z)

	if x < y && x < z && y < z {
		fmt.Print(x, y, z)
	} else if x < y && x < z && z < y {
		fmt.Print(x, z, y)
	} else if y < x && y < z && x < z {
		fmt.Print(y, x, z)
	} else if y < x && y < z && z < y {
		fmt.Print(y, z, x)
	} else if z < x && z < y && x < y {
		fmt.Print(z, x, y)
	} else {
		fmt.Print(z, y, x)
	}

}
