package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d démarre job %d\n", id, job)
        time.Sleep(time.Second)  // Simuler du travail
        fmt.Printf("Worker %d termine job %d\n", id, job)
        results <- job * 2
    }
}

func main() {
    const numJobs = 5
    const numWorkers = 3
    
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    
    // Démarrer les workers
    for w := 1; w <= numWorkers; w++ {
        go worker(w, jobs, results)
    }
    
    // Envoyer les jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collecter et afficher les résultats
    for a := 1; a <= numJobs; a++ {
        res := <-results       // Lire un résultat depuis le channel
        fmt.Println("Résultat :", res)
    }

}