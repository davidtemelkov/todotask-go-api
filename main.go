package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	setupDB()

	router := gin.Default()

	router.GET("/people", getPeople)
	router.POST("/people", createPeople)
	router.PUT("/people", editPeople)
	router.DELETE("/people", deletePeople)

	router.GET("/tasks", getToDoTasks)
	router.POST("/tasks", createTasks)
	router.PUT("/tasks", editTasks)
	router.DELETE("/tasks", deleteTasks)

	router.Run("localhost:8080")
}
