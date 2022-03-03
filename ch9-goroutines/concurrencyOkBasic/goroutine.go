package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go showGoroutine(i, wg)

		//quitando la palabra "go" lo hace secuencial
		//Goroutine #0 with 81ms
		//Goroutine #1 with 87ms
		//Goroutine #9 with 0ms
		//505ms

		//dejando la goroutine se demora menos pero es aleatorio
		//Goroutine #9 with 81ms
		//Goroutine #3 with 87ms
		//Goroutine #0 with 47ms
		//98ms
	}

	wg.Wait()
	duration := time.Since(start).Milliseconds()

	fmt.Printf("%dms\n", duration)
}

func showGoroutine(id int, wg *sync.WaitGroup) {
	delay := rand.Intn(100)
	fmt.Printf("Goroutine #%d with %dms\n", id, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}
