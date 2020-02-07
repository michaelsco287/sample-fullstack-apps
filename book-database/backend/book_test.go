package backend

import (
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
		   "publishers":[
			  {
				 "name":"Simon & Schuster"
			  }
		   ],
		   "title":"Frederick Douglass: Prophet of Freedom",
		   "cover":{
			  "large":"https://covers.openlibrary.org/b/id/8312265-L.jpg",
		   },
		   "authors":[
			{
			   "url":"https://openlibrary.org/authors/OL953055A/David_W._Blight",
			   "name":"David W. Blight"
			}
		 ],
		   "publish_date":"2018"
		}
	 }`))
	ISBN := "9781416590316"
	got := FindBookByISBN("9781416590316")
	want := Book{
		title:         "Frederick Douglass: Prophet of Freedom",
		authors:       []string{"David W. Blight"},
		publishers:    []string{"Simon & Schuster"},
		datePublished: "2018",
		coverURL:      "https://covers.openlibrary.org/b/id/8312265-L.jpg",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindBookByISBN(\"%s\"), expected: %#v but got: %#v", ISBN, want, got)
	}
}
