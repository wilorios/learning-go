//THIS IS AN EXAMPLE THAT SOLVER THE PROBLEM OF "DATA RACE"
package main

import (
	"concurrencyOk/data"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go readBookOK(i, wg, m)
	}

	wg.Wait()
	duration := time.Since(start).Milliseconds()

	fmt.Printf("%dms\n", duration)

	//pero al ejecutarlo con:
	//go run --race .\main.go
	//Finished book: Padre Rico Padre Pobre
	//Finished book: El toque de midas
	//Finished book: Secretos de una mente millonaria
	//Finished book: Retire Rich
	//93ms

}

func readBookOK(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
	data.FinishBookOK(id, m)
	delay := rand.Intn(100)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}
