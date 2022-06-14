package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{Id: "1", Title: "nut", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodos(context *gin.Context) {
	var newTodos todo

	if err := context.BindJSON(&newTodos); err != nil {
		return
	}

	todos = append(todos, newTodos)
	context.IndentedJSON(http.StatusCreated, newTodos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodos)
	router.Run("localhost:8000")
}
