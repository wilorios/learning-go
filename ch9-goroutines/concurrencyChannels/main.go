package main

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

func main() {
	//wait group
	wg := &sync.WaitGroup{}
	//creacion canal(default bidireccional envian y reciben datos)
	IDSChan := make(chan string)
	FakeIDSChan := make(chan string)
	closedChan := make(chan int)

	//cuantas goroutines vamos a esperar
	wg.Add(3)

	go generateIDS(wg, IDSChan, closedChan)
	go generateFakeIDS(wg, FakeIDSChan, closedChan)
	go logIDs(wg, IDSChan, FakeIDSChan, closedChan)

	wg.Wait()
}

//goroutine 1
//se marca el "chan<-" para indicar que el canal es receive-only
func generateIDS(wg *sync.WaitGroup, idsChan chan<- string, closedChan chan<- int) {

	for i := 0; i < 10; i++ {
		id := uuid.New()
		//destino <- origen
		idsChan <- fmt.Sprintf("%d. %s ", i+1, id.String())
	}

	//OJO DEBEMOS CERRAR EL CANAL receive-only
	close(idsChan)
	closedChan <- 1

	wg.Done()
}

//goroutine 2
//se marca el "chan<-" para indicar que el canal es receive-only
func generateFakeIDS(wg *sync.WaitGroup, fakeIdsChan chan<- string, closedChan chan<- int) {

	for i := 0; i < 5; i++ {
		id := uuid.New()
		//destino <- origen
		fakeIdsChan <- fmt.Sprintf("%d. %s ", i+1, id.String())
	}

	//OJO DEBEMOS CERRAR EL CANAL receive-only
	close(fakeIdsChan)
	closedChan <- 1

	wg.Done()
}

//goroutine 3
//se marca el "<-chan" para indicar que el canal es send-only
//para el closedChan el chan es bidireccional porque extrae data y hace un close
func logIDs(wg *sync.WaitGroup, idsChan <-chan string, fakeIdsChan <-chan string, closedChan chan int) {

	closedCounter := 0
	for {
		select {
		case id, ok := <-idsChan:
			if ok {
				fmt.Println("ID: ", id)
			}
		case id, ok := <-fakeIdsChan:
			if ok {
				fmt.Println("FAKE ID: ", id)
			}
		case count, ok := <-closedChan:
			if ok {
				closedCounter += count
			}
		}
		if closedCounter == 2 {
			close(closedChan)
			break
		}
	}
	wg.Done()
}
