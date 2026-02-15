package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)
    
    // Goroutine 1
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "Message de c1"
    }()
    
    // Goroutine 2
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "Message de c2"
    }()
    
    // Select attend le premier channel disponible
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("Reçu:", msg1)
        case msg2 := <-c2:
            fmt.Println("Reçu:", msg2)
        }
    }
    
    // Select avec default (non bloquant)
    messages := make(chan string)
    signals := make(chan bool)

    go func() {
        time.Sleep(2 * time.Second)
        messages <- "Message"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        signals <- true
    }()
    
    time.Sleep(2 * time.Second)

    select {
    case msg := <-messages:
        fmt.Println("Message:", msg)
    default:
        fmt.Println("Aucun message disponible")
    }
    
    // Select avec timeout
    select {
    case msg := <-messages:
        fmt.Println(msg)
    case sig := <-signals:
        fmt.Println(sig)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout après 1 seconde")
    }
}