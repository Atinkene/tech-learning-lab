package main

import (
    "fmt"
    "sync"
    "time"
)

type Cache struct {
    data map[string]string
    mu   sync.RWMutex  // Permet plusieurs lecteurs simultanés
}

func (c *Cache) Lire(key string) (string, bool) {
    c.mu.RLock()           // Verrou lecture (partagé)
    defer c.mu.RUnlock()
    val, ok := c.data[key]
    return val, ok
}

func (c *Cache) Ecrire(key, value string) {
    c.mu.Lock()            // Verrou écriture (exclusif)
    defer c.mu.Unlock()
    c.data[key] = value
}

func main() {
    cache := &Cache{
        data: make(map[string]string),
    }
    
    // Écriture initiale
    cache.Ecrire("ville", "Dakar")
    cache.Ecrire("pays", "Sénégal")
    
    var wg sync.WaitGroup
    
    // 100 lecteurs concurrents (pas de blocage entre eux)
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            if val, ok := cache.Lire("ville"); ok {
                fmt.Printf("Lecteur %d: %s\n", id, val)
            }
            time.Sleep(10 * time.Millisecond)
        }(i)
    }
    
    wg.Wait()
}