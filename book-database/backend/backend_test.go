package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	mockhttp "github.com/karupanerura/go-mock-http-response"
	"github.com/stretchr/testify/assert"
)

func mockResponse(statusCode int, headers map[string]string, body []byte) {
	http.DefaultClient = mockhttp.NewResponseMock(statusCode, headers, body).MakeClient()

}
func TestPostOwnedBooksRoute_ValidISBN(t *testing.T) {
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
			   "url":"https://openlibrary.org/Authors/OL953055A/David_W._Blight",
			   "name":"David W. Blight"
			}
		 ],
		   "publish_date":"2018"
		}
	 }`))

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
	mockResponse(http.StatusOK, map[string]string{"Content-Type": "application/json"}, []byte(`{
		"ISBN:9781416590316":{
		   "publishers":[
			  {
				 "name":"Simon and Schuster"
			  }
		   ],
		   "title":"Frederick Douglass",
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
	 }`))
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/book?isbn=9781416590316", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"Title\":\"Frederick Douglass\",\"Authors\":[\"David W. Blight\"],\"Publishers\":[\"Simon and Schuster\"],\"DatePublished\":\"2018\",\"CoverURL\":\"https://covers.openlibrary.org/b/id/8312265-L.jpg\"}\n", w.Body.String())
}
func TestGetBook_InvalidISBN(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/book?isbn=9782416590316", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "Invalid ISBN (9782416590316) provided.", w.Body.String())
}
