package main

// Créer une calculatrice qui:
// - Demande deux nombres et une opération (+, -, *, /)
// - Affiche le résultat
// - Gère les erreurs (division par zéro, opération invalide)

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func Calculatrice() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez le premier nombre: ")
	reader.Scan()
	num1Str := reader.Text()
	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		fmt.Println("Erreur: Entrée invalide.")
		return
	}

	fmt.Print("Entrez le deuxième nombre: ")
	reader.Scan()
	num2Str := reader.Text()
	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		fmt.Println("Erreur: Entrée invalide.")
		return
	}

	fmt.Print("Entrez l'opération (+, -, *, /): ")
	reader.Scan()
	op := reader.Text()

	var resultat float64
	switch op {
	case "+":
		resultat = num1 + num2
	case "-":
		resultat = num1 - num2
	case "*":
		resultat = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Erreur: Division par zéro.")
			return
		}
		resultat = num1 / num2
	default:
		fmt.Println("Erreur: Opération invalide.")
		return
	}

	fmt.Printf("Le résultat est: %f\n", resultat)
}

func main() {
	Calculatrice()
}