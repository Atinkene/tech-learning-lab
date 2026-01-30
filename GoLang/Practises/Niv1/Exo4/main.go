package main

// Classique FizzBuzz:
// - Nombres 1 à 100
// - Multiples de 3: "Fizz"
// - Multiples de 5: "Buzz"
// - Multiples des deux: "FizzBuzz"

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}