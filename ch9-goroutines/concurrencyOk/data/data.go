package data

import (
	"fmt"
	"sync"
)

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{1, "El toque de midas", false},
	{2, "Padre Rico Padre Pobre", false},
	{3, "Secretos de una mente millonaria", false},
	{4, "Retire Rich", false},
	{5, "Business school", false},
}

func findBookOK(id int, m *sync.RWMutex) (int, *Book) {
	index := -1
	var book *Book

	m.RLock() //AGREGAMOS ESTO PARA EVITAR EL DATA RACE ---> LECTURA
	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
		}
	}
	m.RUnlock() //AGREGAMOS ESTO PARA EVITAR EL DATA RACE, ---> LECTURA
	return index, book
}

func FinishBookOK(id int, m *sync.RWMutex) {
	i, book := findBookOK(id, m)
	if i < 0 {
		return
	}
	m.Lock() //AGREGAMOS ESTO PARA EVITAR EL DATA RACE, --->ESCRITURA
	book.Finished = true
	books[i] = book
	m.Unlock() //AGREGAMOS ESTO PARA EVITAR EL DATA RACE, --->ESCRITURA
	fmt.Printf("Finished book: %s\n", book.Title)
}
