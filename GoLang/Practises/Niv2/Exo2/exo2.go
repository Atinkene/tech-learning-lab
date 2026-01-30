package main

import "fmt"

func main() {
    ch := make(chan int)
    numGoroutines := 1000

    for i := 1; i <= numGoroutines; i++ {
        go func(i int) {
            ch <- i
            fmt.Printf("L'id %d est envoyé\n", i)
        }(i)
    }

    fmt.Println("Les ids reçus sont :")
    for i := 0; i < numGoroutines; i++ {
        id := <-ch
        fmt.Println(id)
    }
}
