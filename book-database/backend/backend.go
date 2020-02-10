package backend

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/ownedbooks", func(c *gin.Context) {
		lib := Library{}
		isbn := c.PostForm("isbn")
		foundBook, err := FindBookByISBN(isbn)
		if err != nil {
			c.String(400, fmt.Sprintf("Invalid ISBN (%v) provided.", isbn))
			return
		}
		lib.OwnBook(foundBook)
		c.String(200, fmt.Sprintf("Added \"%s\" to your list of owned books.", foundBook.title))
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
