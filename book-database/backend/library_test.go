package backend

import (
	"reflect"
	"testing"
)

func TestAddBook(t *testing.T) {
	lib := Library{}
	book := Book{
		title: "Test",
	}
	lib.AddBook(book)
	got := lib.booksOwned
	want := []Book{book}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("AddBook(\"%v\"), expected: %#v but got: %#v", book, want, got)
	}
}
