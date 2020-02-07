package backend

type Library struct {
	booksOwned []Book
}

func (l *Library) AddBook(b Book) {
	l.booksOwned = append(l.booksOwned, b)
}
