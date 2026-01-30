package main

import "fmt"

func main() {
    nom := "Massina"
    age := 25
    moyenne := 15.75
    
    // Print (pas de retour ligne)
    fmt.Print("Hello ")
    fmt.Print("World\n")
    
    // Println (avec retour ligne)
    fmt.Println("Bonjour", nom)
    
    // Printf (formatage)
    fmt.Printf("Nom: %s, Age: %d ans\n", nom, age)
    fmt.Printf("Moyenne: %.2f/20\n", moyenne)
    
    // Verbes de formatage courants
    fmt.Printf("%v\n", nom)         // Valeur par défaut
    fmt.Printf("%+v\n", struct{X int}{42})  // Struct avec noms de champs
    fmt.Printf("%#v\n", []int{1,2,3})       // Syntaxe Go
    fmt.Printf("%T\n", age)         // Type
    fmt.Printf("%t\n", true)        // Boolean
    fmt.Printf("%d\n", 42)          // Décimal
    fmt.Printf("%b\n", 42)          // Binaire
    fmt.Printf("%o\n", 42)          // Octal
    fmt.Printf("%x\n", 42)          // Hexadécimal
    fmt.Printf("%f\n", 3.14159)     // Float
    fmt.Printf("%.2f\n", 3.14159)   // Float avec 2 décimales
    fmt.Printf("%s\n", "texte")     // String
    fmt.Printf("%q\n", "texte")     // String quotée
    fmt.Printf("%p\n", &age)        // Pointeur (adresse)
    
    // Sprintf (retourne string au lieu d'afficher)
    message := fmt.Sprintf("Bonjour %s, tu as %d ans", nom, age)
    fmt.Println(message)
    
    // Errorf (créer une erreur formatée)
    err := fmt.Errorf("erreur de connexion à %s:%d", "localhost", 8080)
    fmt.Println(err)
    
    // Scan (lire depuis stdin)
    var prenom string
    fmt.Print("Entrez votre prénom: ")
    fmt.Scan(&prenom)
    fmt.Println("Bonjour", prenom)
}