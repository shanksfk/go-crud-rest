package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Book struct to hold book information
type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

// Initialize a slice to hold books
var books = []Book{
    {ID: "1", Title: "The Go Programming Language", Author: "Alan A. A. Donovan and Brian W. Kernighan"},
    {ID: "2", Title: "Go in Action", Author: "William Kennedy"},
}

//GET
func GetBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

//POST
func PostBook(c *gin.Context) {
    var newBook Book
    if err := c.BindJSON(&newBook); err != nil {
        return
    }
    books = append(books, newBook)
    c.IndentedJSON(http.StatusCreated, newBook)
}

//GET
func GetBookByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range books {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// UPDATE
func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var updatedBook Book
    if err := c.BindJSON(&updatedBook); err != nil {
        return
    }
    for i, a := range books {
        if a.ID == id {
            books[i] = updatedBook
            c.IndentedJSON(http.StatusOK, updatedBook)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// DELETE
func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    for i, a := range books {
        if a.ID == id {
            books = append(books[:i], books[i+1:]...)
            c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
