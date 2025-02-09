package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}


var books = []book{

	{Id: "1", Title: "The screct", Author: "Bainda", Price: 210.76},
	{Id: "2", Title: "The Happines", Author: "Vijay", Price: 310.76},
	{Id: "3", Title: "Greatful", Author: "Vijay", Price: 410.10},
}

func getallbook(a *gin.Context) {

	a.IndentedJSON(http.StatusOK, books)
}

func postbook(a *gin.Context) {

	var newbook book

	a.BindJSON(&newbook) // user given data format is json BindJSON(&newbook)will  convert go struct in server

	books = append(books, newbook)

	a.IndentedJSON(http.StatusCreated, newbook) // saying frontend or client new data susscessfully added.

}

func getbook(a *gin.Context) {

	var id = a.Param("id")

	for _, b := range books {

		if b.Id == id {

			a.IndentedJSON(http.StatusOK, b)
			return
		}
	}

	a.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Book not found"})
}

func main() {

	r := gin.Default()
	r.GET("/books", getallbook)
	r.POST("/books", postbook)
	r.GET("/books/:id", getbook)

	r.Run(":8080")
}

/*
1. first step

 go mod init my - creating go mod file
 rm go.mod      - delet the the mod file

2. secound step

go mod tidy  - this will create the sum file and creating the dependance package.

*/
