package main

import (
	"fmt"
	"sync"
	"time"
)

// Fonction lourde à paralléliser
func traiterDonnee(id int, data int) int {
	time.Sleep(100 * time.Millisecond)
	return data * data
}

// Fan-Out: distribuer le travail à plusieurs goroutines
func fanOut(in <-chan int, numWorkers int) []<-chan int {
	channels := make([]<-chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		workerID := i
		out := make(chan int)
		channels[i] = out

		go func() {
			for data := range in {
				result := traiterDonnee(workerID, data)
				out <- result
			}
			close(out)
		}()
	}

	return channels
}

// Fan-In: combiner les résultats de plusieurs channels
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// Source de données
	input := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()

	// Fan-Out vers 3 workers
	workers := fanOut(input, 3)

	// Fan-In des résultats
	results := fanIn(workers...)

	// Consommer les résultats
	for result := range results {
		fmt.Println(result)
	}
}