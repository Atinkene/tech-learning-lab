package main

import (
    "fmt"
    "time"
)

func main() {
    // Créer un channel
    messages := make(chan string)
    
    // Goroutine qui envoie des données
    go func() {
        messages <- "Premier message"
        messages <- "Deuxième message"
        messages <- "Troisième message"
    }()
    
    // Recevoir des données (bloquant)
    msg1 := <-messages
    msg2 := <-messages
    msg3 := <-messages
    
    fmt.Println(msg1)
    fmt.Println(msg2)
    fmt.Println(msg3)
    
    // Channel avec timeout
    reponse := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        reponse <- "Réponse après 2 secondes"
    }()
    
    // Attendre max 1 seconde
    select {
    case msg := <-reponse:
        fmt.Println(msg)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout dépassé!")
    }
}