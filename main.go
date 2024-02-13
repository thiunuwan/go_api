package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Qty    int    `json:"qty"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Qty: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Qty: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Qty: 6},
}


func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK,books)
}

func createBook(c *gin.Context){
	var newBook book

	if err := c.BindJSON(&newBook); err != nil{
     	return 
	}
	books = append(books,newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}


func main(){
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
    router.Run("localhost:8080")
}