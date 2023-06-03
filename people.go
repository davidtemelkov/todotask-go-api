package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Person struct {
	ID   uint
	NAME string
	AGE  uint
}

func getPeople(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		var people []Person
		result := db.Find(&people)

		if result.Error != nil {
			panic(result.Error)
		}

		c.IndentedJSON(http.StatusOK, people)

		return
	}

	var person Person
	result := db.First(&person, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found."})

			return
		}

		panic(result.Error)
	}

	c.IndentedJSON(http.StatusOK, person)
}
