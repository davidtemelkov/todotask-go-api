package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	setupDB()

	router := gin.Default()
	router.GET("/people", getPeople)
	router.GET("/tasks", getToDoTasks)
	router.Run("localhost:8080")
}
