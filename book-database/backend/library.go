package backend

type Library struct {
	booksOwned []Book
}

func (l *Library) OwnBook(b Book) {
	l.booksOwned = append(l.booksOwned, b)
}
