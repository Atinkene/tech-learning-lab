package main

import (
	"fmt"
	"math/rand/v2"
	"time"
	"sync"
)

func (turnOver *turnOver)orderProcessing (order Order) float64 {
	
	time.Sleep(200 * time.Millisecond)
	finalAmount := order.amount * 1.18
	turnOver.mu.Lock()
	turnOver.turnOverValue += finalAmount
	turnOver.mu.Unlock()
	return finalAmount

}

type Order struct {
	ID int
	amount float64
}

type turnOver struct {
	turnOverValue float64
	mu sync.Mutex
}

func (turnOver *turnOver)workerPool (

	id int , 
	orders <- chan Order, 
	results chan <- float64, 
	wg *sync.WaitGroup,

) {

	defer wg.Done()

	for order := range orders {

		fmt.Printf("Worker %d start order %d\n", id, order.ID)
		res := turnOver.orderProcessing(order)
		results <- res
		fmt.Printf("Worker %d end order %d\n", id, order.ID)

	}
}

func main () {

	startAt := time.Now()

	defer func () {
		endAt := time.Now()
		fmt.Printf("Total time taken : %v\n", endAt.Sub(startAt))
	}()

	numWorkers := 3
	numOrders := 10
	turnOver := &turnOver{
		turnOverValue : rand.Float64() * 1000000000,
	}

	var wg sync.WaitGroup
	orders := make(chan Order, numOrders)
	results := make(chan float64, numOrders)

	go func () {
		for i := 0; i < numOrders; i++ {
			orders <- Order {
				ID : rand.IntN(100),
				amount : rand.Float64() * 10000000000000,
			}
		}
		close(orders)
	}()

	for i := 0 ; i < numWorkers ; i++ {
		wg.Add(1)
		go turnOver.workerPool(i, orders, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	
	timeOut := time.After(1 * time.Second)
	for {
		select {
			case res, ok := <- results : 
				if !ok {
					fmt.Println("Turnover : ", turnOver.turnOverValue)
					return
				}
				fmt.Println("end of process", res)
			
			case <- timeOut:
				fmt.Println("Cannot finish processing on time")
				return
		}
	}

}