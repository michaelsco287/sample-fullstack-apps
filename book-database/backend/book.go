package backend

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

type Book struct {
	title         string
	authors       []string
	publishers    []string
	datePublished string
	coverURL      string
}

func FindBookByISBN(ISBN string) Book {
	URLToRequest := fmt.Sprintf("https://openlibrary.org/api/books?bibkeys=ISBN:%s&jscmd=data&format=json", ISBN)
	resp, err := http.Get(URLToRequest)
	if err != nil {
		fmt.Printf("Unexpected error contacting OpenLibrary API")
	}
	requestedBook := Book{}
	body, _ := ioutil.ReadAll(resp.Body)
	JSONResponseToBook(string(body), &requestedBook)
	defer resp.Body.Close()

	return requestedBook
}

func JSONResponseToBook(body string, book *Book) {
	fmt.Print(gjson.Get(body, "ISBN*.authors.#.name").Array())
	book.title = gjson.Get(body, "ISBN*.title").String()
	book.authors = []string{gjson.Get(body, "ISBN*.authors.#.name").Array()[0].String()}
	book.publishers = []string{gjson.Get(body, "ISBN*.publishers.#.name").Array()[0].String()}
	book.datePublished = gjson.Get(body, "ISBN*.publish_date").String()
	book.coverURL = gjson.Get(body, "ISBN*.cover.large").String()
}
