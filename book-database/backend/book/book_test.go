package book

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	mockhttp "github.com/karupanerura/go-mock-http-response"
)

func mockResponse(statusCode int, headers map[string]string, body []byte) {
	http.DefaultClient = mockhttp.NewResponseMock(statusCode, headers, body).MakeClient()

}

func TestFindBookByISBN(t *testing.T) {
	mockResponse(http.StatusOK, map[string]string{"Content-Type": "application/json"}, []byte(`{
		"ISBN:9781416590316":{
		   "Publishers":[
			  {
				 "name":"Simon & Schuster"
			  }
		   ],
		   "Title":"Frederick Douglass: Prophet of Freedom",
		   "cover":{
			  "large":"https://covers.openlibrary.org/b/id/8312265-L.jpg",
		   },
		   "Authors":[
			{
			   "url":"https://openlibrary.org/Authors/OL953055A/David_W._Blight",
			   "name":"David W. Blight"
			}
		 ],
		   "publish_date":"2018"
		}
	 }`))
	ISBN := "9781416590316"
	got, _ := FindBookByISBN(ISBN)
	want := Book{
		Title:         "Frederick Douglass: Prophet of Freedom",
		Authors:       []string{"David W. Blight"},
		Publishers:    []string{"Simon & Schuster"},
		DatePublished: "2018",
		CoverURL:      "https://covers.openlibrary.org/b/id/8312265-L.jpg",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindBookByISBN(\"%s\"), expected: %#v but got: %#v", ISBN, want, got)
	}
}

func TestFindBookByISBN_WhenISBNLengthNotTenOrThirteen(t *testing.T) {
	ISBN := "12321"
	_, got := FindBookByISBN(ISBN)
	want := errors.New("You did not provide a valid ISBN number")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindBookByISBN(\"%s\"), expected: %#v but got: %#v", ISBN, want, got)
	}
}
func TestFindBookByISBN_WhenInvalidISBN(t *testing.T) {
	ISBN := "9780136091817"
	_, got := FindBookByISBN(ISBN)
	want := errors.New("You did not provide a valid ISBN number")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindBookByISBN(\"%s\"), expected: %#v but got: %#v", ISBN, want, got)
	}
}
