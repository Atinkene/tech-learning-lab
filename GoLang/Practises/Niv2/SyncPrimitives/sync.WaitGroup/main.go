package main

import (
    "fmt"
    "sync"
    "time"
)

func tache(id int, wg *sync.WaitGroup) {
    defer wg.Done()  // Décrémenter le compteur à la fin
    
    fmt.Printf("Tâche %d démarre\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Tâche %d termine\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    // Lancer 5 goroutines
    for i := 1; i <= 5; i++ {
        wg.Add(1)  // Incrémenter le compteur
        go tache(i, &wg)
    }
    
    fmt.Println("Attente de toutes les tâches...")
    wg.Wait()  // Bloquer jusqu'à ce que le compteur soit 0
    fmt.Println("Toutes les tâches terminées")
}