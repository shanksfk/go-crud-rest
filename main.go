
package main

import (
    "github.com/gin-gonic/gin"
	"fmt"
)

func main() {
    router := gin.Default()

    // Using the exported handler functions
    router.GET("/books", GetBooks)
    router.POST("/books", PostBook)
    router.GET("/books/:id", GetBookByID)
    router.PUT("/books/:id", UpdateBook)
    router.DELETE("/books/:id", DeleteBook)
	fmt.Println("Running the server...")
    router.Run(":8080")
}
