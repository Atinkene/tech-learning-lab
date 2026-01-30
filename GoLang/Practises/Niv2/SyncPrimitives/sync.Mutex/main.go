package main

import (
    "fmt"
    "sync"
)

type CompteBancaire struct {
    solde float64
    mu    sync.Mutex  // Protège solde
}

func (c *CompteBancaire) Crediter(montant float64) {
    c.mu.Lock()           // Verrouiller
    defer c.mu.Unlock()   // Déverrouiller à la fin
    c.solde += montant
}

func (c *CompteBancaire) Debiter(montant float64) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.solde -= montant
}

func (c *CompteBancaire) Solde() float64 {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.solde
}

func main() {
    compte := &CompteBancaire{solde: 1000}
    var wg sync.WaitGroup
    
    // 100 crédits concurrents
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            compte.Crediter(10)
        }()
    }
    
    // 50 débits concurrents
    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            compte.Debiter(5)
        }()
    }
    
    wg.Wait()
    fmt.Printf("Solde final: %.2f\n", compte.Solde())
}