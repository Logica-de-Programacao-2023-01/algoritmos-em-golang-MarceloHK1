package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else if i%15 == 0 {
			fmt.Print("FizzBuzz")
		} else {
			fmt.Print(i)
		}
	}
}
