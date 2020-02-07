package backend

type Library struct {
	booksOwned  []Book
	booksWanted []Book
}

func (l *Library) OwnBook(b Book) {
	l.booksOwned = append(l.booksOwned, b)
}
func (l *Library) WantBook(b Book) {
	l.booksWanted = append(l.booksWanted, b)
}
