package main

import (
	"net/http"
	"time"

	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Todotask struct {
	ID          uint
	NAME        string
	DESCRIPTION string `gorm:"column:DESCRIPTION"`
	CREATED_ON  time.Time
	DONE        bool
	CREATOR_ID  uint `gorm:"column:CREATOR_ID"`
	ASSIGNEE_ID uint `gorm:"column:ASSIGNEE_ID"`
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

func createTasks(c *gin.Context) {
	var newTask Todotask

	if err := c.BindJSON(&newTask); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&newTask).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newTask)
}
