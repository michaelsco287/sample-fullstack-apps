package backend

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostOwnedBooksRoute(t *testing.T) {
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

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/ownedbooks", bytes.NewBufferString("isbn=9781416590316"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Added \"Frederick Douglass: Prophet of Freedom\" to your list of owned books.", w.Body.String())
}