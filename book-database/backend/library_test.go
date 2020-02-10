package backend

import (
	"reflect"
	"testing"
)

func TestOwnBook(t *testing.T) {
	lib := Library{}
	book := Book{
		Title: "Test",
	}
	lib.OwnBook(book)
	got := lib.booksOwned
	want := []Book{book}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("OwnBook(\"%v\"), expected: %#v but got: %#v", book, want, got)
	}
}

func TestWantBook(t *testing.T) {
	lib := Library{}
	book := Book{
		Title: "Fave Book",
	}
	lib.WantBook(book)
	got := lib.booksWanted
	want := []Book{book}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WantBook(\"%v\"), expected: %#v but got: %#v", book, want, got)
	}
}
