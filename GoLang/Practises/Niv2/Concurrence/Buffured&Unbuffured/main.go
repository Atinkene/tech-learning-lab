package main

import "fmt"

func main() {
    // Unbuffered channel (bloquant)
    ch1 := make(chan int)
    
    go func() {
        ch1 <- 42  // Bloque jusqu'à ce que quelqu'un reçoive
    }()
    
    valeur := <-ch1  // Bloque jusqu'à ce que quelqu'un envoie
    fmt.Println(valeur)
    
    
    // Buffered channel (non bloquant jusqu'à capacité max)
    ch2 := make(chan int, 3)  // Buffer de taille 3
    
    // On peut envoyer 3 valeurs sans bloquer
    ch2 <- 1
    ch2 <- 2
    ch2 <- 3
    // ch2 <- 4  // Bloquerait car buffer plein
    
    // Recevoir
    fmt.Println(<-ch2)  // 1
    fmt.Println(<-ch2)  // 2
    fmt.Println(<-ch2)  // 3

    
    // Fermer un channel
    nombres := make(chan int, 5)
    
    go func() {
        for i := 0; i < 10; i++ {
            nombres <- i
        }
        close(nombres)  // Signal que plus de données arrivent
    }()
    fmt.Println("Starting")
     nombres <- 10
    // Recevoir jusqu'à fermeture
    for num := range nombres {
        fmt.Println(num)
    }
    
    // // Vérifier si un channel est fermé
    // valeur, ok := <-nombres
    // if !ok {
    //     fmt.Println("Channel fermé")
    // }
}