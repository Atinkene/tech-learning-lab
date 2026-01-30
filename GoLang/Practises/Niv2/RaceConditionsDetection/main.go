package main

import (
    "fmt"
    "sync"
)

// MAUVAIS: Race condition
func exempleRace() {
    compteur := 0
    var wg sync.WaitGroup
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            compteur++  // ❌ Accès concurrent non protégé
        }()
    }
    
    wg.Wait()
    fmt.Println("Compteur (race):", compteur)  // Résultat imprévisible
}

// BON: Avec protection
func exempleSafe() {
    compteur := 0
    var wg sync.WaitGroup
    var mu sync.Mutex
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            compteur++  // ✅ Protégé
            mu.Unlock()
        }()
    }
    
    wg.Wait()
    fmt.Println("Compteur (safe):", compteur)  // Toujours 1000
}

func main() {
    exempleRace()
    exempleSafe()
}

// Exécuter avec: go run -race main.go
// Le flag -race détecte les race conditions