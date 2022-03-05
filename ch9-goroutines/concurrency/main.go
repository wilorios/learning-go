package main

import (
	"concurrency/data"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go readBookWithDataRace(i, wg)
	}

	wg.Wait()
	duration := time.Since(start).Milliseconds()

	fmt.Printf("%dms\n", duration)

	//al ejecutarlo pareciese todo ok

	//go run .\main.go
	//Finished book: Retire Rich
	//Finished book: El toque de midas
	//Finished book: Padre Rico Padre Pobre
	//Finished book: Secretos de una mente millonaria
	//89ms

	//pero al ejecutarlo con:
	//go run --race .\main.go
	//Finished book: Secretos de una mente millonaria
	//Finished book: Retire Rich
	//Finished book: Padre Rico Padre Pobre
	//WARNING: DATA RACE
	//Read at 0x000000ff6910 by goroutine 8:
	//main/data.findBook()
	//D:/Perfiles/wrios/workspace-go/concurrency/data/data.go:23 +0x68
	//Previous write at 0x000000ff6910 by goroutine 10:
	//166ms
	//Found 1 data race(s)
	//exit status 66
}

func readBookWithDataRace(id int, wg *sync.WaitGroup) {
	data.FinishBookWithDataRace(id)
	delay := rand.Intn(100)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}
