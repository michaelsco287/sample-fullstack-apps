package library

import (
	"github.com/NicksonT/sample-fullstack-apps/book-database/backend/book"
)

type Library struct {
	BooksOwned  []book.Book
	BooksWanted []book.Book
}

func (l *Library) OwnBook(b book.Book) {
	l.BooksOwned = append(l.BooksOwned, b)
}
func (l *Library) WantBook(b book.Book) {
	l.BooksWanted = append(l.BooksWanted, b)
}
