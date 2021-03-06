package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	mockhttp "github.com/karupanerura/go-mock-http-response"
	"github.com/stretchr/testify/assert"
)

var openAPIresponse = []byte(`{
	"ISBN:9781416590316":{
	   "publishers":[
		  {
			 "name":"Simon & Schuster"
		  }
	   ],
	   "identifiers": {"isbn_13": ["9781416590316"], "lccn": ["2018007511"], "openlibrary": ["OL26624981M"], "isbn_10": ["1416590315"]}
	   "title":"Frederick Douglass: Prophet of Freedom",
	   "cover":{
		  "large":"https://covers.openlibrary.org/b/id/8312265-L.jpg",
	   },
	   "authors":[
		{
		   "url":"https://openlibrary.org/Authors/OL953055A/David_W._Blight",
		   "name":"David W. Blight"
		}
	 ],
	   "publish_date":"2018"
	}
 }`)

func mockResponse(statusCode int, headers map[string]string, body []byte) {
	http.DefaultClient = mockhttp.NewResponseMock(statusCode, headers, body).MakeClient()

}
func TestPostOwnedBooksRoute_ValidISBN(t *testing.T) {
	mockResponse(http.StatusOK, map[string]string{"Content-Type": "application/json"}, openAPIresponse)

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/ownedbooks", bytes.NewBufferString("isbn=9781416590316"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Added \"Frederick Douglass: Prophet of Freedom\" to your list of owned books.", w.Body.String())
}
func TestPostOwnedBooksRoute_InvalidISBN(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/ownedbooks", bytes.NewBufferString("isbn=9782416590316"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "Invalid ISBN (9782416590316) provided.", w.Body.String())
}

func TestGetBook_ValidISBN(t *testing.T) {
	mockResponse(http.StatusOK, map[string]string{"Content-Type": "application/json"}, openAPIresponse)
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/book?isbn=9781416590316", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"Title\":\"Frederick Douglass: Prophet of Freedom\",\"Authors\":[\"David W. Blight\"],\"Publishers\":[\"Simon \\u0026 Schuster\"],\"DatePublished\":\"2018\",\"CoverURL\":\"https://covers.openlibrary.org/b/id/8312265-L.jpg\",\"ISBN10\":\"1416590315\",\"ISBN13\":\"9781416590316\"}\n", w.Body.String())
}
func TestGetBook_InvalidISBN(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/book?isbn=9782416590316", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "Invalid ISBN (9782416590316) provided.", w.Body.String())
}
func TestPostWantedBooksRoute_ValidISBN(t *testing.T) {
	mockResponse(http.StatusOK, map[string]string{"Content-Type": "application/json"}, openAPIresponse)

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/wantedbooks", bytes.NewBufferString("isbn=9781416590316"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Added \"Frederick Douglass: Prophet of Freedom\" to your list of wanted books.", w.Body.String())
}
func TestPostWantedBooksRoute_InvalidISBN(t *testing.T) {

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/wantedbooks", bytes.NewBufferString("isbn=9782416590316"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "Invalid ISBN (9782416590316) provided.", w.Body.String())
}
func TestGetOwnedBooks(t *testing.T) {
	router := setupRouter()
	wGet := httptest.NewRecorder()
	wPost := httptest.NewRecorder()
	mockResponse(http.StatusOK, map[string]string{"Content-Type": "application/json"}, openAPIresponse)
	getReq, _ := http.NewRequest("GET", "/ownedbooks", nil)
	postReq, _ := http.NewRequest("POST", "/ownedbooks", bytes.NewBufferString("isbn=9781416590316"))
	postReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(wPost, postReq)
	router.ServeHTTP(wGet, getReq)

	assert.Equal(t, 200, wGet.Code)
	assert.Equal(t, "[{\"Title\":\"Frederick Douglass: Prophet of Freedom\",\"Authors\":[\"David W. Blight\"],\"Publishers\":[\"Simon \\u0026 Schuster\"],\"DatePublished\":\"2018\",\"CoverURL\":\"https://covers.openlibrary.org/b/id/8312265-L.jpg\",\"ISBN10\":\"1416590315\",\"ISBN13\":\"9781416590316\"}]\n", wGet.Body.String())
}

func TestGetWantedBooks(t *testing.T) {
	router := setupRouter()
	wGet := httptest.NewRecorder()
	wPost := httptest.NewRecorder()
	mockResponse(http.StatusOK, map[string]string{"Content-Type": "application/json"}, openAPIresponse)
	getReq, _ := http.NewRequest("GET", "/wantedbooks", nil)
	postReq, _ := http.NewRequest("POST", "/wantedbooks", bytes.NewBufferString("isbn=9781416590316"))
	postReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(wPost, postReq)
	router.ServeHTTP(wGet, getReq)

	assert.Equal(t, 200, wGet.Code)
	assert.Equal(t, "[{\"Title\":\"Frederick Douglass: Prophet of Freedom\",\"Authors\":[\"David W. Blight\"],\"Publishers\":[\"Simon \\u0026 Schuster\"],\"DatePublished\":\"2018\",\"CoverURL\":\"https://covers.openlibrary.org/b/id/8312265-L.jpg\",\"ISBN10\":\"1416590315\",\"ISBN13\":\"9781416590316\"}]\n", wGet.Body.String())
}
