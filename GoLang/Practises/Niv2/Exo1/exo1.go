package main

import (
	"fmt"
	"sync"
)

func afficherId(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Goroutine %d démarre\n", id)
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 3

	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)
		go afficherId(i, &wg)
	}
	wg.Wait()	
}