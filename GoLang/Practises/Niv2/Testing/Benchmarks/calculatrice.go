package calculatrice

import (
	"fmt"
)

func Additionner(a, b int) int {
    return a + b
}

func Soustraire(a, b int) int {
    return a - b
}

func Multiplier(a, b int) int {
    return a * b
}

func Diviser(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division par zéro")
    }
    return a / b, nil
}