package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{

	{Id: "1", Item: "clean room", Completed: false},
	{Id: "2", Item: "Read book", Completed: false},
	{Id: "3", Item: "Record video", Completed: false},
}

func getTodos(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {

	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil { // err variable stands for to catch the error

		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)

}

func getTodo(context *gin.Context) {

	id := context.Param("id")

	todo, err := getTodobyid(id)

	if err != nil {

		context.IndentedJSON(http.StatusFound, gin.H{"Message": "Todo not found "})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodostatus(context *gin.Context) {

	id := context.Param("id")

	todo, err := getTodobyid(id)

	if err != nil {

		context.IndentedJSON(http.StatusFound, gin.H{"Message": "Todo not found "})
		return
	}

	//todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodobyid(id string) (*todo, error) {

	for i, t := range todos {

		if t.Id == id {

			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found ")
}

func main() {

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodostatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")

}

/*

1. context.IndentedJSON(http.StatusOK, todos)

context.IndentedJSON(http.StatusOK, todos), the purpose is to send a JSON response back to the client.

http.StatusOK: This is an HTTP status code constant. In this case, it represents 200,
which indicates a successful response. It tells the client that the request was successful.

code is responsible for formatting the todos data as JSON and sending it back to the
client in a way that it can be understood and used by the client's application or frontend.


2. context.BindJSON(&newTodo)
in post man your adding the data in json format.just creating the new entry and sending to the server.

When you use context.BindJSON(&newTodo), you're telling the server to take the JSON data sent in the request body and convert it into a Go struct (newTodo).
This allows you to work with the data in a structured and typed manner within your Go code.

if you send JSON data like {"name": "John Doe", "age": 30}, context.BindJSON(&newTodo) will parse it and
populate the name and age fields of the newTodo struct. This makes it much easier to work with the data
in your Go code, as you can now access it using dot notation like newTodo.name and newTodo.age.

context.IndentedJSON(http.StatusCreated, newTodo), you're telling the client that a new todo item has been successfully created,
and you're providing the details of the newly created todo item in the response body. This is commonly used in APIs that handle resource creation.

In numerical terms, http.StatusCreated is equivalent to 201.
When a server responds with StatusCreated, it typically includes a Location header in the response to indicate the URL of the newly created resource.

context.IndentedJSON - this one always using for sending an HTTP response back to the client.

3.getTodo

context.IndentedJSON(http.StatusOK, todo) is telling the client that the request to retrieve the todo item was successful,
and it is providing the todo item as the response in JSON format. The client will receive an HTTP response with a status code of 200 OK along with the JSON representation of the todo item.

Purpose

while using Rest API GET and POST we are mentioned http.StatusOK and http.StatusCreated
this and and all using for main purpose just telling to fronted and client.


*/
