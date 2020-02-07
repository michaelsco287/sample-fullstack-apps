package backend

import (
	"reflect"
	"testing"
)

func TestOwnBook(t *testing.T) {
	lib := Library{}
	book := Book{
		title: "Test",
	}
	lib.OwnBook(book)
	got := lib.booksOwned
	want := []Book{book}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("OwnBook(\"%v\"), expected: %#v but got: %#v", book, want, got)
	}
}
