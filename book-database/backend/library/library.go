package library

import (
	"github.com/NicksonT/sample-fullstack-apps/book-database/backend/book"
)

type Library struct {
	booksOwned  []book.Book
	booksWanted []book.Book
}

func (l *Library) OwnBook(b book.Book) {
	l.booksOwned = append(l.booksOwned, b)
}
func (l *Library) WantBook(b book.Book) {
	l.booksWanted = append(l.booksWanted, b)
}
