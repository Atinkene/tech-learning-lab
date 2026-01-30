package main

import (
	"fmt"
	"time"
    "sync"
)
func traiterDonnee(workerId, data int) int {
    fmt.Printf("Worker %d traite la donnée %d\n", workerId, data)
    time.Sleep(100 * time.Millisecond) 
    fmt.Printf("Worker %d a terminé le traitement de la donnée %d\n", workerId, data)
    return data * data
}

func fanOut(in <-chan int, numWorkers int) []<-chan int {
	channels := make([]<-chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		out := make(chan int)
		channels[i] = out

		go func() {
			for data := range in {
				result := traiterDonnee(i, data)
				out <- result
			}
			close(out)
		}()
	}

	return channels
}

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
    jobs := make(chan int)

    go func () {
        for i := 1; i <= 20; i++ {
            jobs <- i
        }
        close(jobs)
    }()

    nombreWorkers := 4

    workers := fanOut(jobs, nombreWorkers)
    resultat := fanIn(workers...)

    for res := range resultat {
        fmt.Println("Résultat reçu :", res)
    }

}
