package backend

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

type Book struct {
	title         string
	authors       []string
	publishers    []string
	datePublished string
	coverURL      string
}

func FindBookByISBN(ISBN string) (Book, error) {
	if !isValidISBNFormat(ISBN) {
		return Book{}, errors.New("You did not provide a valid ISBN number")
	}
	URLToRequest := fmt.Sprintf("https://openlibrary.org/api/books?bibkeys=ISBN:%s&jscmd=data&format=json", ISBN)
	resp, err := http.Get(URLToRequest)
	if err != nil {
		fmt.Printf("Unexpected error contacting OpenLibrary API")
	}
	requestedBook := Book{}
	body, _ := ioutil.ReadAll(resp.Body)
	responseToBook(string(body), &requestedBook)
	defer resp.Body.Close()

	return requestedBook, nil
}

func isValidISBNFormat(isbn string) bool {
	isbn = strings.Replace(strings.TrimSpace(isbn), "-", "", -1)
	if len(isbn) != 13 && len(isbn) != 10 { // valid ISBN are 10 or 13 digits
		return false
	}
	if len(isbn) == 10 {
		total := 0
		for pos, num := range Reverse(isbn) {
			total += int(num) * (pos + 1)
		}
		if total%11 != 0 {
			return false
		}
	}
	if len(isbn) == 13 {

		total := 0

		for pos, num := range isbn {
			if (pos+1)%2 == 0 {
				total += int(num) * 3
			} else {
				total += int(num)
			}
		}
		fmt.Print(total)
		if total%10 != 0 {
			return false
		}
	}
	return true
}

func responseToBook(body string, book *Book) {
	fmt.Print(gjson.Get(body, "ISBN*.authors.#.name").Array())
	book.title = gjson.Get(body, "ISBN*.title").String()
	book.authors = []string{gjson.Get(body, "ISBN*.authors.#.name").Array()[0].String()}
	book.publishers = []string{gjson.Get(body, "ISBN*.publishers.#.name").Array()[0].String()}
	book.datePublished = gjson.Get(body, "ISBN*.publish_date").String()
	book.coverURL = gjson.Get(body, "ISBN*.cover.large").String()
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
