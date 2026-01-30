package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    // Variables d'environnement
    home := os.Getenv("HOME")
    fmt.Println("HOME:", home)
    
    os.Setenv("MA_VARIABLE", "valeur")
    fmt.Println(os.Getenv("MA_VARIABLE"))
    
    // Arguments de ligne de commande
    fmt.Println("Nom du programme:", os.Args[0])
    if len(os.Args) > 1 {
        fmt.Println("Arguments:", os.Args[1:])
    }
    
    // Créer un fichier
    fichier, err := os.Create("test.txt")
    if err != nil {
        fmt.Println("Erreur:", err)
        return
    }
    defer fichier.Close()
    
    // Écrire dans le fichier
    fichier.WriteString("Bonjour Dakar!\n")
    fichier.WriteString("Go est génial!\n")
    
    // Lire un fichier (tout d'un coup)
    contenu, err := os.ReadFile("test.txt")
    if err != nil {
        fmt.Println("Erreur:", err)
        return
    }
    fmt.Println(string(contenu))
    
    // Écrire un fichier (tout d'un coup)
    data := []byte("Nouvelle ligne\n")
    os.WriteFile("nouveau.txt", data, 0644)
    
    // Ouvrir un fichier existant
    f, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("Erreur:", err)
        return
    }
    defer f.Close()
    
    // Lire avec io
    buffer := make([]byte, 1024)
    n, err := f.Read(buffer)
    if err != nil && err != io.EOF {
        fmt.Println("Erreur:", err)
        return
    }
    fmt.Printf("Lu %d octets\n", n)
    
    // Informations sur fichier
    info, _ := os.Stat("test.txt")
    fmt.Println("Nom:", info.Name())
    fmt.Println("Taille:", info.Size())
    fmt.Println("Permissions:", info.Mode())
    fmt.Println("Modifié:", info.ModTime())
    fmt.Println("Est un répertoire?", info.IsDir())
    
    // Vérifier l'existence
    if _, err := os.Stat("inexistant.txt"); os.IsNotExist(err) {
        fmt.Println("Le fichier n'existe pas")
    }
    
    // Supprimer un fichier
    os.Remove("nouveau.txt")
    
    // Créer un répertoire
    os.Mkdir("monrep", 0755)
    os.MkdirAll("rep/sous-rep/sous-sous", 0755)  // Récursif
    dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		} else {
			fmt.Println("Repertoire ", dir)
	} 
    // Lister un répertoire
    entries, _ := os.ReadDir(".")
    for _, entry := range entries {
        fmt.Println(entry.Name(), entry.IsDir())
    }
    
    // Nettoyer
    os.Remove("test.txt")
    os.RemoveAll("monrep")
    os.RemoveAll("rep")
}