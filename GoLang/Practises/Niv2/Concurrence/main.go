package main

import (
    "fmt"
    "time"
)

// func direBonjour(s string) {
//     fmt.Println("Bonjour ", s)
//     time.Sleep(1 * time.Second)

// }

// func main() {
//     go direBonjour("Massina")
//     go direBonjour("Amadou")    
//     time.Sleep(2 * time.Second)
// }

func main() {
    m := make(chan string)

    go func(m chan string) {
        m <- "Salut depuis la goroutine 0"
        m <- "Salut depuis la goroutine 1"
        m <- "Salut depuis la goroutine 2"
    }(m)
    
    m1 := <- m
    m2 := <- m
    m3 := <- m

    fmt.Println(m1)
    fmt.Println(m2)
    fmt.Println(m3)

    select {
    case m := <- m:
        fmt.Println("Reçu:", m)
    case <- time.After(1 * time.Second):
        fmt.Println("Timeout atteint, aucune donnée reçue")
    }

}