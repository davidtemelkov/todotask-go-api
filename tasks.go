package main

import (
	"net/http"
	"time"

	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Todotask struct {
	ID           uint
	NAME         string
	DESCRIPTION  string
	CREATED_ON   time.Time
	DONE         bool
	ASSIGNMENTID uint
}

func getToDoTasks(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		var todoTasks []Todotask
		result := db.Find(&todoTasks)

		if result.Error != nil {
			panic(result.Error)
		}

		c.IndentedJSON(http.StatusOK, todoTasks)
		return
	}

	var toDoTask Todotask
	result := db.First(&toDoTask, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found."})
			return
		}

		panic(result.Error)
	}

	c.IndentedJSON(http.StatusOK, toDoTask)
}
