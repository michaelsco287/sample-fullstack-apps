package main

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
		c.String(200, fmt.Sprintf("Added \"%s\" to your list of owned books.", foundBook.Title))
	})
	r.GET("/book", func(c *gin.Context) {
		isbn := c.Query("isbn")
		foundBook, err := FindBookByISBN(isbn)
		if err != nil {
			c.String(400, fmt.Sprintf("Invalid ISBN (%v) provided.", isbn))
			return
		}
		if err != nil {
			c.String(500, "Issue with server, please try again later")
			return
		}
		fmt.Print(foundBook)
		c.JSON(200, foundBook)
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
