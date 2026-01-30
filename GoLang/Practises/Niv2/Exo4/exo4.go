package main

import (
	"fmt"
	"time"
)

func main() {
    data := make(chan string)
    quit := make(chan bool)

    go func() {
        time.Sleep(200 * time.Millisecond)
        data <- "job terminé"
    }()
    go func() {
        quit <- true
    }()

    select {
        case message := <-data:
            fmt.Println("Reçu:", message)
        case <-time.After(100 * time.Millisecond):
            fmt.Println("Timeout atteint, aucune donnée reçue")
    }
}
