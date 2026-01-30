package main

import (
	"fmt"
)
func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func main() {

	contact := map[string]string{
		"nom":        "",
		"prénom":     "",
		"adresse":    "",
		"téléphone":  "",
		"email":      "",
	}	

	contact["nom"] = "Doe"
	contact["prénom"] = "John"
	contact["adresse"] = "123 Rue Principale"
	contact["téléphone"] = "0123456789"
	contact["email"] = "john.doe@example.com"


	contact["postal"] = "11500"
	printMap(contact)

	fmt.Println("----- Après recherche du téléphone -----")
	fmt.Printf("Téléphone: %s \n", contact["téléphone"])

	fmt.Printf("apres recherche \n")



	fmt.Println("----- Après suppression du champ postal -----")
	delete(contact, "postal")
	printMap(contact)

}