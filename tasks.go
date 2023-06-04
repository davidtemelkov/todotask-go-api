package main

import (
	"net/http"
	"time"

	"errors"

	"strconv"

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

	uintNum, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	uintId := uint(uintNum)

	task, err := getTaskByID(uintId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
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

func editTasks(c *gin.Context) {
	var newTask Todotask

	if err := c.BindJSON(&newTask); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task Todotask
	task, err := getTaskByID(newTask.ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	task.NAME = newTask.NAME
	task.DESCRIPTION = newTask.DESCRIPTION
	task.CREATED_ON = newTask.CREATED_ON
	task.DONE = newTask.DONE
	task.CREATOR_ID = newTask.CREATOR_ID
	task.ASSIGNEE_ID = newTask.ASSIGNEE_ID

	if err := db.Save(&task).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func getTaskByID(id uint) (Todotask, error) {
	var task Todotask
	result := db.First(&task, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return task, errors.New("Task not found")
		}
		return task, result.Error
	}
	return task, nil
}
