package main

import (
	"fmt"
)

func main() {
	nbK := 10
	nbJ := 10
	nbLigne := 10

	for i := 0; i < nbLigne; i++ {
		for j := 0; j < nbJ; j++ {
			for k := 0; k < nbK; k++ {
				fmt.Printf("%d x %d = %d\t", k, j, k*j)
			}
			fmt.Println()
		}
		fmt.Println()
	}
}