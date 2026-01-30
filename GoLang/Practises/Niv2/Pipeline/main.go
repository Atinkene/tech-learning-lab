package main

import "fmt"

// Étape 1: Générer des nombres
func generer(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// Étape 2: Multiplier par 2
func doubler(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * 2
        }
        close(out)
    }()
    return out
}

// Étape 3: Ajouter 10
func ajouter10(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n + 10
        }
        close(out)
    }()
    return out
}

func main() {
    // Pipeline: generer → doubler → ajouter10
    c := generer(1, 2, 3, 4, 5)
    c = doubler(c)
    c = ajouter10(c)
    
    // Consommer les résultats
    for resultat := range c {
        fmt.Println(resultat)
    }
}