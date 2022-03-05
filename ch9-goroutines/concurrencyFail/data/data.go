package data

import (
	"fmt"
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

func findBookWithDataRace(id int) (int, *Book) {
	index := -1
	var book *Book

	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
		}
	}
	return index, book
}

func FinishBookWithDataRace(id int) {
	i, book := findBookWithDataRace(id)
	if i < 0 {
		return
	}
	book.Finished = true
	books[i] = book
	fmt.Printf("Finished book: %s\n", book.Title)
}
