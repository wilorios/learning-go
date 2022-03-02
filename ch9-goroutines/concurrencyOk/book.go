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
	{1, "El toque de midas", true},
	{2, "Padre Rico Padre Pobre", true},
}

func findBookOK(id int, m *sync.RWMutex) (int, *Book) {
	index := -1
	var book *Book
	m.RLock() //agregamos esto para evitar el data race
	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
		}
	}
	m.RUnlock() //agregamos esto para evitar el data race
	return index, book
}

func FinishBookOK(id int, m *sync.RWMutex) {
	i, book := findBookOK(id, m)
	if i < 0 {
		return
	}
	m.Lock() //agregamos esto para evitar el data race
	book.Finished = true
	books[i] = book
	m.Unlock() //agregamos esto para evitar el data race
	fmt.Printf("Finished book: %s\n ", book.Title)
}
